package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
	"unsafe"

	"github.com/creack/pty"
	"github.com/gliderlabs/ssh"
)

func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

func main() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, "Connect to qgb/sshd \n")
		scmd:="/bin/bash"
		_, err := os.Stat(scmd)
		if os.IsNotExist(err) {
			scmd="/data/data/com.termux/files/usr/bin/bash"
			_, err := os.Stat(scmd)
			if os.IsNotExist(err) {
				scmd="/system/bin/sh"
			}

		}
		/*
		if _, err := os.Stat(scmd); errors.Is(err, os.ErrNotExist) {
		  // path/to/whatever does not exist
		}*/
		cmd := exec.Command(scmd)
					
		ptyReq, winCh, isPty := s.Pty()
		if isPty {
			cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
			f, err := pty.Start(cmd)
			if err != nil {
				panic(err)
			}
			go func() {
				for win := range winCh {
					setWinsize(f, win.Width, win.Height)
				}
			}()
			go func() {
				io.Copy(f, s) // stdin
			}()
			io.Copy(s, f) // stdout
			cmd.Wait()
		} else {
			io.WriteString(s, "No PTY requested.\n")
			s.Exit(1)
		}
	})

	log.Println("starting ssh server on port 2222... tty.go")
//	log.Fatal(ssh.ListenAndServe(":2222", nil))
        log.Fatal(ssh.ListenAndServe(":2222", nil,
                ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
                        return pass == "qgb"
                }),
        ))
}
