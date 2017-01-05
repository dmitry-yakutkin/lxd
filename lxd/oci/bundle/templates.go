package bundle

const (
	// TODO: make real template
	hostnameTemplateString = `
f77e35c1320f`

	// TODO: make real template
	hostsTemplateString = `
#
# /etc/hosts: static lookup table for host names
#

#<ip-address>	<hostname.domain.org>	<hostname>
127.0.0.1	localhost.localdomain	localhost
::1		localhost.localdomain	localhost
127.0.0.1	f77e35c1320f`

	// TODO: make real template
	resolvconfTemplateString = `
nameserver 8.8.8.8
nameserver 8.8.4.4`

	// TODO: extend with additional data
	configTemplateString = `
{
    "ociVersion": "1.0.0-rc1",
    "platform": {
	"os": "linux",
	"arch": "amd64"
    },
    "process": {
	"terminal": true,
	"user": {
	    "uid": 0,
	    "gid": 0,
	    "additionalGids": [
		10
	    ]
	},
	"args": [
	    "sh"
	],
	"env": [
	    "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
	    "HOSTNAME=f77e35c1320f",
	    "TERM=xterm"
	],
	"cwd": "/",
	"capabilities": [
	    "CAP_CHOWN",
	    "CAP_DAC_OVERRIDE",
	    "CAP_FSETID",
	    "CAP_FOWNER",
	    "CAP_MKNOD",
	    "CAP_NET_RAW",
	    "CAP_SETGID",
	    "CAP_SETUID",
	    "CAP_SETFCAP",
	    "CAP_SETPCAP",
	    "CAP_NET_BIND_SERVICE",
	    "CAP_SYS_CHROOT",
	    "CAP_KILL",
	    "CAP_AUDIT_WRITE"
	]
    },
    "root": {
	"path": "{{.BundlePath}}/rootfs"
    },
    "hostname": "f77e35c1320f",
    "mounts": [
	{
	    "destination": "/proc",
	    "type": "proc",
	    "source": "proc",
	    "options": [
		"nosuid",
		"noexec",
		"nodev"
	    ]
	},
	{
	    "destination": "/dev",
	    "type": "tmpfs",
	    "source": "tmpfs",
	    "options": [
		"nosuid",
		"strictatime",
		"mode=755"
	    ]
	},
	{
	    "destination": "/dev/pts",
	    "type": "devpts",
	    "source": "devpts",
	    "options": [
		"nosuid",
		"noexec",
		"newinstance",
		"ptmxmode=0666",
		"mode=0620",
		"gid=5"
	    ]
	},
	{
	    "destination": "/sys",
	    "type": "sysfs",
	    "source": "sysfs",
	    "options": [
		"nosuid",
		"noexec",
		"nodev",
		"ro"
	    ]
	},
	{
	    "destination": "/sys/fs/cgroup",
	    "type": "cgroup",
	    "source": "cgroup",
	    "options": [
		"ro",
		"nosuid",
		"noexec",
		"nodev"
	    ]
	},
	{
	    "destination": "/dev/mqueue",
	    "type": "mqueue",
	    "source": "mqueue",
	    "options": [
		"nosuid",
		"noexec",
		"nodev"
	    ]
	},
	{
	    "destination": "/etc/resolv.conf",
	    "type": "bind",
	    "source": "{{.BundlePath}}/resolv.conf",
	    "options": [
		"rbind",
		"rprivate"
	    ]
	},
	{
	    "destination": "/etc/hostname",
	    "type": "bind",
	    "source": "{{.BundlePath}}/hostname",
	    "options": [
		"rbind",
		"rprivate"
	    ]
	},
	{
	    "destination": "/etc/hosts",
	    "type": "bind",
	    "source": "{{.BundlePath}}/hosts",
	    "options": [
		"rbind",
		"rprivate"
	    ]
	},
	{
	    "destination": "/dev/shm",
	    "type": "bind",
	    "source": "{{.BundlePath}}/shm",
	    "options": [
		"rbind",
		"rprivate"
	    ]
	}
    ],
    "linux": {
	"resources": {
	    "devices": [
		{
		    "allow": false,
		    "access": "rwm"
		},
		{
		    "allow": true,
		    "type": "c",
		    "major": 1,
		    "minor": 5,
		    "access": "rwm"
		},
		{
		    "allow": true,
		    "type": "c",
		    "major": 1,
		    "minor": 3,
		    "access": "rwm"
		},
		{
		    "allow": true,
		    "type": "c",
		    "major": 1,
		    "minor": 9,
		    "access": "rwm"
		},
		{
		    "allow": true,
		    "type": "c",
		    "major": 1,
		    "minor": 8,
		    "access": "rwm"
		},
		{
		    "allow": true,
		    "type": "c",
		    "major": 5,
		    "minor": 0,
		    "access": "rwm"
		},
		{
		    "allow": true,
		    "type": "c",
		    "major": 5,
		    "minor": 1,
		    "access": "rwm"
		},
		{
		    "allow": false,
		    "type": "c",
		    "major": 10,
		    "minor": 229,
		    "access": "rwm"
		}
	    ],
	    "disableOOMKiller": false,
	    "oomScoreAdj": 0,
	    "memory": {
		"kernelTCP": null,
		"swappiness": 18446744073709551615
	    },
	    "cpu": {},
	    "pids": {
		"limit": 0
	    },
	    "blockIO": {
		"blkioWeight": 0
	    }
	},
	"namespaces": [
	    {
		"type": "mount"
	    },
	    {
		"type": "uts"
	    },
	    {
		"type": "pid"
	    },
	    {
		"type": "ipc"
	    }
	],
	"devices": [
	    {
		"path": "/dev/fuse",
		"type": "c",
		"major": 10,
		"minor": 229,
		"fileMode": 438,
		"uid": 0,
		"gid": 0
	    }
	],
	"seccomp": {
	    "defaultAction": "SCMP_ACT_ERRNO",
	    "architectures": [
		"SCMP_ARCH_X86_64",
		"SCMP_ARCH_X86",
		"SCMP_ARCH_X32"
	    ],
	    "syscalls": [
		{
		    "name": "accept",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "accept4",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "access",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "alarm",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "bind",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "brk",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "capget",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "capset",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "chdir",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "chmod",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "chown",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "chown32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "clock_getres",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "clock_gettime",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "clock_nanosleep",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "close",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "connect",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "copy_file_range",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "creat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "dup",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "dup2",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "dup3",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "epoll_create",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "epoll_create1",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "epoll_ctl",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "epoll_ctl_old",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "epoll_pwait",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "epoll_wait",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "epoll_wait_old",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "eventfd",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "eventfd2",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "execve",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "execveat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "exit",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "exit_group",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "faccessat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fadvise64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fadvise64_64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fallocate",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fanotify_mark",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fchdir",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fchmod",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fchmodat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fchown",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fchown32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fchownat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fcntl",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fcntl64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fdatasync",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fgetxattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "flistxattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "flock",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fork",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fremovexattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fsetxattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fstat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fstat64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fstatat64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fstatfs",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fstatfs64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "fsync",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "ftruncate",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "ftruncate64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "futex",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "futimesat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getcpu",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getcwd",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getdents",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getdents64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getegid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getegid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "geteuid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "geteuid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getgid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getgid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getgroups",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getgroups32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getitimer",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getpeername",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getpgid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getpgrp",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getpid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getppid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getpriority",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getrandom",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getresgid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getresgid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getresuid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getresuid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getrlimit",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "get_robust_list",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getrusage",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getsid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getsockname",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getsockopt",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "get_thread_area",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "gettid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "gettimeofday",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getuid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getuid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "getxattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "inotify_add_watch",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "inotify_init",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "inotify_init1",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "inotify_rm_watch",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "io_cancel",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "ioctl",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "io_destroy",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "io_getevents",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "ioprio_get",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "ioprio_set",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "io_setup",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "io_submit",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "ipc",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "kill",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "lchown",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "lchown32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "lgetxattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "link",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "linkat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "listen",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "listxattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "llistxattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "_llseek",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "lremovexattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "lseek",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "lsetxattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "lstat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "lstat64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "madvise",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "memfd_create",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mincore",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mkdir",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mkdirat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mknod",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mknodat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mlock",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mlock2",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mlockall",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mmap",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mmap2",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mprotect",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mq_getsetattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mq_notify",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mq_open",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mq_timedreceive",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mq_timedsend",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mq_unlink",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "mremap",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "msgctl",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "msgget",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "msgrcv",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "msgsnd",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "msync",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "munlock",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "munlockall",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "munmap",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "nanosleep",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "newfstatat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "_newselect",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "open",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "openat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "pause",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "personality",
		    "action": "SCMP_ACT_ALLOW",
		    "args": [
			{
			    "index": 0,
			    "value": 0,
			    "valueTwo": 0,
			    "op": "SCMP_CMP_EQ"
			}
		    ]
		},
		{
		    "name": "personality",
		    "action": "SCMP_ACT_ALLOW",
		    "args": [
			{
			    "index": 0,
			    "value": 8,
			    "valueTwo": 0,
			    "op": "SCMP_CMP_EQ"
			}
		    ]
		},
		{
		    "name": "personality",
		    "action": "SCMP_ACT_ALLOW",
		    "args": [
			{
			    "index": 0,
			    "value": 4294967295,
			    "valueTwo": 0,
			    "op": "SCMP_CMP_EQ"
			}
		    ]
		},
		{
		    "name": "pipe",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "pipe2",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "poll",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "ppoll",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "prctl",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "pread64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "preadv",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "prlimit64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "pselect6",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "pwrite64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "pwritev",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "read",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "readahead",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "readlink",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "readlinkat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "readv",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "recv",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "recvfrom",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "recvmmsg",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "recvmsg",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "remap_file_pages",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "removexattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "rename",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "renameat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "renameat2",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "restart_syscall",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "rmdir",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "rt_sigaction",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "rt_sigpending",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "rt_sigprocmask",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "rt_sigqueueinfo",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "rt_sigreturn",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "rt_sigsuspend",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "rt_sigtimedwait",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "rt_tgsigqueueinfo",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sched_getaffinity",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sched_getattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sched_getparam",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sched_get_priority_max",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sched_get_priority_min",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sched_getscheduler",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sched_rr_get_interval",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sched_setaffinity",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sched_setattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sched_setparam",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sched_setscheduler",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sched_yield",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "seccomp",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "select",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "semctl",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "semget",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "semop",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "semtimedop",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "send",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sendfile",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sendfile64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sendmmsg",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sendmsg",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sendto",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setfsgid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setfsgid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setfsuid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setfsuid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setgid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setgid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setgroups",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setgroups32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setitimer",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setpgid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setpriority",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setregid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setregid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setresgid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setresgid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setresuid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setresuid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setreuid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setreuid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setrlimit",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "set_robust_list",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setsid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setsockopt",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "set_thread_area",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "set_tid_address",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setuid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setuid32",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "setxattr",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "shmat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "shmctl",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "shmdt",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "shmget",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "shutdown",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sigaltstack",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "signalfd",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "signalfd4",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sigreturn",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "socket",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "socketcall",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "socketpair",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "splice",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "stat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "stat64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "statfs",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "statfs64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "symlink",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "symlinkat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sync",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sync_file_range",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "syncfs",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "sysinfo",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "syslog",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "tee",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "tgkill",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "time",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "timer_create",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "timer_delete",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "timerfd_create",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "timerfd_gettime",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "timerfd_settime",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "timer_getoverrun",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "timer_gettime",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "timer_settime",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "times",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "tkill",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "truncate",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "truncate64",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "ugetrlimit",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "umask",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "uname",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "unlink",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "unlinkat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "utime",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "utimensat",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "utimes",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "vfork",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "vmsplice",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "wait4",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "waitid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "waitpid",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "write",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "writev",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "arch_prctl",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "modify_ldt",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "chroot",
		    "action": "SCMP_ACT_ALLOW"
		},
		{
		    "name": "clone",
		    "action": "SCMP_ACT_ALLOW",
		    "args": [
			{
			    "index": 0,
			    "value": 2080505856,
			    "valueTwo": 0,
			    "op": "SCMP_CMP_MASKED_EQ"
			}
		    ]
		}
	    ]
	},
	"maskedPaths": [
	    "/proc/kcore",
	    "/proc/latency_stats",
	    "/proc/timer_list",
	    "/proc/timer_stats",
	    "/proc/sched_debug"
	],
	"readonlyPaths": [
	    "/proc/asound",
	    "/proc/bus",
	    "/proc/fs",
	    "/proc/irq",
	    "/proc/sys",
	    "/proc/sysrq-trigger"
	]
    },
    "solaris": {
	"cappedCPU": {},
	"cappedMemory": {}
    }
}`
)
