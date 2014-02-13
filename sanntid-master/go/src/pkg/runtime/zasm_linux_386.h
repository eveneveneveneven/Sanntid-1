// auto generated by go tool dist

// On Linux systems, what we call 0(GS) and 4(GS) for g and m
// turn into %gs:-8 and %gs:-4 (using gcc syntax to denote
// what the machine sees as opposed to 8l input).
// 8l rewrites 0(GS) and 4(GS) into these.
//
// On Linux Xen, it is not allowed to use %gs:-8 and %gs:-4
// directly.  Instead, we have to store %gs:0 into a temporary
// register and then use -8(%reg) and -4(%reg).  This kind
// of addressing is correct even when not running Xen.
//
// 8l can rewrite MOVL 0(GS), CX into the appropriate pair
// of mov instructions, using CX as the intermediate register
// (safe because CX is about to be written to anyway).
// But 8l cannot handle other instructions, like storing into 0(GS),
// which is where these macros come into play.
// get_tls sets up the temporary and then g and r use it.
//
// The final wrinkle is that get_tls needs to read from %gs:0,
// but in 8l input it's called 8(GS), because 8l is going to
// subtract 8 from all the offsets, as described above.
#define	get_tls(r)	MOVL 8(GS), r
#define	g(r)	-8(r)
#define	m(r)	-4(r)
#define gobuf_sp 0
#define gobuf_pc 4
#define gobuf_g 8
#define g_stackguard 0
#define g_stackbase 4
#define g_defer 8
#define g_panic 12
#define g_sched 16
#define g_gcstack 28
#define g_gcsp 32
#define g_gcguard 36
#define g_stack0 40
#define g_entry 44
#define g_alllink 48
#define g_param 52
#define g_status 56
#define g_goid 60
#define g_selgen 64
#define g_waitreason 68
#define g_schedlink 72
#define g_readyonstop 76
#define g_ispanic 77
#define g_m 80
#define g_lockedm 84
#define g_idlem 88
#define g_sig 92
#define g_writenbuf 96
#define g_writebuf 100
#define g_sigcode0 104
#define g_sigcode1 108
#define g_sigpc 112
#define g_gopc 116
#define g_end 120
#define m_g0 0
#define m_morepc 4
#define m_moreargp 8
#define m_morebuf 12
#define m_moreframesize 24
#define m_moreargsize 28
#define m_cret 32
#define m_procid 36
#define m_gsignal 44
#define m_tls 48
#define m_curg 80
#define m_id 84
#define m_mallocing 88
#define m_gcing 92
#define m_locks 96
#define m_nomemprof 100
#define m_waitnextg 104
#define m_dying 108
#define m_profilehz 112
#define m_helpgc 116
#define m_fastrand 120
#define m_ncgocall 124
#define m_havenextg 132
#define m_nextg 136
#define m_alllink 140
#define m_schedlink 144
#define m_machport 148
#define m_mcache 152
#define m_stackalloc 156
#define m_lockedg 160
#define m_idleg 164
#define m_createstack 168
#define m_freglo 296
#define m_freghi 360
#define m_fflag 424
#define m_nextwaitm 428
#define m_waitsema 432
#define m_waitsemacount 436
#define m_waitsemalock 440
#define m_end 444
#define wincall_fn 0
#define wincall_n 4
#define wincall_args 8
#define wincall_r1 12
#define wincall_r2 16
#define wincall_err 20
