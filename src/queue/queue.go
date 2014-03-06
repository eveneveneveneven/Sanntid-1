package main

import(
	"fmt"
	"json"
	"network"
)

type(
	MoveDir int
	OrderDir int
)

const(
	N_BUTTONS int = 3
	N_FLOORS  int = 4
	MAX_ORDERS int = 10
	N_ELEVATORS int = 2

	ORDER_UP OrderDir = iota //matched with FLOOR for actuall order
	ORDER_DOWN
	ORDER_INTERNAL

	MOVE_UP MoveDir = iota //defines av elevators direction 
	MOVE_DOWN
	MOVE_STOP
)

type Elevator struct{
	//constant
	Ip string
	//subject to change (will trigger select)
	orderQueue []Order
	direction MoveDir
	lastFloor int
}

type Order struct{
	floor int
	orientation OrderDir
}

type Cost struct{
	cost int
	order Order
	ip string
}


func GetElevatorEligibilityScore(elevator Elevator, order Order) int {
	score := 0
	//correct direction plays 4 points difference
	if elevator.direction == MOVE_STOP{
		score += 0
	}else if ((elevator.direction == MOVE_UP) && (order.floor > elevator.lastFloor)) || ((elevator.direction == MOVE_DOWN) && (order.floor < elevator.lastFloor)){ //hvis bestilling er i riktig retning
		score += 4
	}else{
		score -= 4
	}
	// each order in queue before this order plays 1 point (NOTE: the internal and both the external orders play part consequently)
	score -= GetNumberOfStopsBeforeOrder(elevator, order)
	return score
}

func GetNumberOfStopsBeforeOrder(elevator Elevator, order Order)int{
	placement := GetInsertOrderPlacement(elevator, order)
	stops := placement
	//fmt.Println("GetNumberOfStopsBeforeOrder: placement == ", placement)
	for j:= 0; j < placement; j++{ //Removing common objective orders from score
		fmt.Println(elevator.orderQueue[j].floor == elevator.orderQueue[j+1].floor)
		fmt.Println(elevator.orderQueue[j].orientation == ORDER_INTERNAL || elevator.orderQueue[j+1].orientation == ORDER_INTERNAL)
		if (elevator.orderQueue[j].floor == elevator.orderQueue[j+1].floor) && (elevator.orderQueue[j].orientation == ORDER_INTERNAL || elevator.orderQueue[j+1].orientation == ORDER_INTERNAL){
			fmt.Println(stops)
			j += 1
			stops -= 1
		}
	}
	return stops
}

func GetInsertOrderPlacement(elevator Elevator, order Order) int{
	priOrder := GetInsertOrderPriority(elevator, order)
	//fmt.Println("GetInserOrderPriority : order == ", order, "priOrder == ", priOrder)
	for i := 0; i < MAX_ORDERS; i++{
		if elevator.orderQueue[i] == order{
			fmt.Println("ERROR in InsertOrder: identical order in queue")
			break
		}else if (GetInsertOrderPriority(elevator, elevator.orderQueue[i]) >= priOrder) && (elevator.orderQueue[i].floor >= order.floor) {
			return i
		}
	}
	return -1
}

func GetInsertOrderPriority(elevator Elevator, order Order) int{
		if order.floor == 0{
			fmt.Println("WARNING GetInsertOrderPriority: order.floor == 0")
			return 5
		}else if elevator.direction == MOVE_UP{
			if order.floor > elevator.lastFloor{
				if order.orientation == ORDER_UP || order.orientation == ORDER_INTERNAL{
					return 1
				}else if order.orientation == ORDER_DOWN{
					return 2
				}
			}else if order.floor <= elevator.lastFloor{
				if order.orientation == ORDER_DOWN || order.orientation == ORDER_INTERNAL{
					return 3
				}else if order.orientation == ORDER_UP{
					return 4
				}
			}
		}else if order.orientation == ORDER_DOWN{
			if order.floor < elevator.lastFloor{
				if order.orientation == ORDER_DOWN || order.orientation == ORDER_INTERNAL{
					return 1
				}else if order.orientation == ORDER_UP{
					return 2
				}
			}else if order.floor >= elevator.lastFloor{
				if order.orientation == ORDER_UP || order.orientation == ORDER_INTERNAL{
					return 3
				}else if order.orientation == ORDER_DOWN{
					return 4
				}
			}
		}
		return -1
}

func InsertOrder(elevator Elevator, order Order){
	if order.floor == 0{
		fmt.Println("ERROR in InsertOrder: order.floor == 0")
		return
	}
	placement := GetInsertOrderPlacement(elevator, order)
	if placement == -1{
		fmt.Println("WARNING in InsertOrder: order existing, insertion cancelled")
	}
	var temp, insert Order
	insert = order
	for i := placement; i <MAX_ORDERS; i++{
		temp = elevator.orderQueue[i]
		elevator.orderQueue[i] = insert
		insert = temp
	}
}

func GetLocalElevatorIndex(elevators []Elevator)int{
	localIp := network.GetLocalIp()
	for i := 0; i < N_ELEVATORS; i++{
		if elevators[i].Ip == localIp{
			return i
		}
	}
}

//func HandleNewOrderBidding(){}

//Update Global
func HandleDeadElev(elevators []Elevator, deadIp string){}



func main(){
	//Making situation picture
	localElevatorIndex := GetLocalElevatorIndex
	fmt.Println("LocalElevatorIndex of ", elevators[localElevatorIndex].Ip, " = ", LocalElevatorIndex)

	elevators := make([]Elevator, N_ELEVATORS)
	

	//Channels
	elevatorsChan := make(chan []Elevator)
	newOrdersChan := make(chan Order, 6*N_ELEVATORS)
	localOrdersChan := make(chan Order, 6)
	receivedCostChan := make(chan Cost, MAX_ORDERS)

	var newOrder, localOrder Order
	var receivedCost, localCost Cost

	//go network.ListenForNewOrders(newOrdersChan)
	//go driver.ListenForLocalOrders(localOrdersChan)
	//go network.ListenForElevatorsUpdate(elevatorsChan)
	//go network.ListenForCostList(costChan)

	for select{
		case elevators = <-elevatorsChan:
		case localOrder = <- localOrdersChan:
			InsertOrder(elevators[localElevatorIndex], localOrder)
			network.SendElevatorsUpdate(elevators)
		case newOrder = <-newOrdersChan: //recieves new order and replies with sending local cost
			localCost = Cost{GetElevatorEligibilityScore(elevator, newOrder), newOrder, elevators[LocalElevatorIndex].Ip)
			network.SendCost(localCost)
		case receivedCost = <- receivedCostChan:

	}

	//SendElevatorsUpdate
	//SendCost

}