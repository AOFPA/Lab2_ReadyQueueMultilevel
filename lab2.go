package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	cpu1   string
	cpu2   string
	ready1 []string
	ready2 []string
	ready3 []string
	ready4 []string
	io1    []string
	io2    []string
	io3    []string
	io4    []string

	r1 int
	r2 int
	r3 int

	cpu1p   string
	cpu2p   string
	ready1p []string
	ready2p []string
	ready3p []string
	ready4p []string
	io1p    []string
	io2p    []string
	io3p    []string
	io4p    []string
)

func initx() {
	cpu1 = ""
	cpu2 = ""
	ready1 = make([]string, 10)
	ready2 = make([]string, 10)
	ready3 = make([]string, 10)
	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)

	r1 = 0
	r2 = 0
	r3 = 0

	cpu1p = ""
	cpu2p = ""
	ready1p = make([]string, 10)
	ready2p = make([]string, 10)
	ready3p = make([]string, 10)
	io1p = make([]string, 10)
	io2p = make([]string, 10)
	io3p = make([]string, 10)
	io4p = make([]string, 10)
}

func showProcess() {
	fmt.Printf("\n****Dev By CHOKCHAI JAMNOI****\n")
	fmt.Printf("CPU[1]>%s \n", cpu1)
	fmt.Printf("CPU[2]>%s \n", cpu2)
	fmt.Printf("Ready[1]>")
	for i := range ready1 {
		fmt.Printf("%s|", ready1[i])
	}
	fmt.Printf("\nReady[2]>")
	for i := range ready2 {
		fmt.Printf("%s|", ready2[i])
	}
	fmt.Printf("\nReady[3]>")
	for i := range ready3 {
		fmt.Printf("%s|", ready3[i])
	}
	fmt.Println()
	fmt.Printf("I/O[1]>")
	for i := range io1 {
		fmt.Printf("%s|", io1[i])
	}
	fmt.Println()
	fmt.Printf("I/O[2]>")
	for i := range io2 {
		fmt.Printf("%s|", io2[i])
	}
	fmt.Println()
	fmt.Printf("I/O[3]>")
	for i := range io3 {
		fmt.Printf("%s|", io3[i])
	}
	fmt.Println()
	fmt.Printf("I/O[4]>")
	for i := range io4 {
		fmt.Printf("%s|", io4[i])
	}
	fmt.Println()
	fmt.Printf("r1=%d|r2=%d|r3=%d|", r1, r2, r3)
	fmt.Printf("\ncommand>")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func insertQ(q []string, data string, qplt []string, plt string) {
	for i := range q {
		if q[i] == "" {
			q[i] = data
			qplt[i] = plt
			break
		}
	}
}

func newProcess(p string, plt string) {
	if cpu1 == "" {
		cpu1 = p
		cpu1p = plt
		addPriority(plt)
	} else if cpu2 == "" {
		cpu2 = p
		cpu2p = plt
		addPriority(plt)
	} else {
		if plt == "1" {
			insertQ(ready1, p, ready1p, plt)
		} else if plt == "2" {
			insertQ(ready2, p, ready2p, plt)
		} else if plt == "3" {
			insertQ(ready3, p, ready3p, plt)
		}
	}
}

func addPriority(p string) {
	if p == "1" {
		r1++
	} else if p == "2" {
		r2++
	} else if p == "3" {
		r3++
	}
}

func checkPriority() {
	if r1 == 3 {
		r1 = 0
		if r2 == 3 {
			r2 = 0
		}
	} else if r2 == 3 {
		r2 = 0
	} else if r3 == 3 {
		r3 = 0
	}
}

func terminate(cpuName string) {
	if cpuName == "cpu1" {
		if r1 < 3 && ready1[0] != "" {
			cpu1, cpu1p = deleteQ(ready1, ready1p)
		} else if r2 < 3 && ready2[0] != "" {
			cpu1, cpu1p = deleteQ(ready2, ready2p)
		} else if r3 < 3 && ready3[0] != "" {
			cpu1, cpu1p = deleteQ(ready3, ready3p)
		} else if ready1[0] == "" && ready2[0] == "" && ready3[0] == "" {
			cpu1 = ""
			cpu1p = ""
		}
		checkPriority()
		addPriority(cpu1p)
	} else if cpuName == "cpu2" {
		if r1 < 3 && ready1[0] != "" {
			cpu2, cpu2p = deleteQ(ready1, ready1p)
		} else if r2 < 3 && ready2[0] != "" {
			cpu2, cpu2p = deleteQ(ready2, ready2p)
		} else if r3 < 3 && ready3[0] != "" {
			cpu2, cpu2p = deleteQ(ready3, ready3p)
		} else if ready1[0] == "" && ready2[0] == "" && ready3[0] == "" {
			cpu2 = ""
			cpu2p = ""
		}
		checkPriority()
		addPriority(cpu2p)
	}
}
func deleteQ(q []string, plt []string) (string, string) {
	result := q[0]
	resultp := plt[0]
	for i := range q {
		if i == 0 {
			continue
		}
		q[i-1] = q[i]
		plt[i-1] = plt[i]
	}
	q[9] = ""
	plt[9] = ""
	return result, resultp
}

func expire(cpuName string) {
	if cpuName == "cpu1" {
		plt := cpu1p
		if plt == "1" {
			insertQ(ready1, cpu1, ready1p, cpu1p)
		} else if plt == "2" {
			insertQ(ready2, cpu1, ready2p, cpu1p)
		} else if plt == "3" {
			insertQ(ready3, cpu1, ready3p, cpu1p)
		}
	} else if cpuName == "cpu2" {
		plt := cpu2p
		if plt == "1" {
			insertQ(ready1, cpu2, ready1p, cpu2p)
		} else if plt == "2" {
			insertQ(ready2, cpu2, ready2p, cpu2p)
		} else if plt == "3" {
			insertQ(ready3, cpu2, ready3p, cpu2p)
		}
	}
	nextQ := ""
	nextPlt := ""
	if r1 < 3 && ready1[0] != "" {
		nextQ, nextPlt = deleteQ(ready1, ready1p)
	} else if r2 < 3 && ready2[0] != "" {
		nextQ, nextPlt = deleteQ(ready2, ready2p)
		if r2 < 2 {
			checkPriority()
		}
	} else if r3 < 3 && ready3[0] != "" {
		nextQ, nextPlt = deleteQ(ready3, ready3p)
		checkPriority()
	}
	addPriority(nextPlt)

	if nextQ == "" {
		return
	}

	if cpuName == "cpu1" {
		cpu1 = nextQ
		cpu1p = nextPlt
	} else if cpuName == "cpu2" {
		cpu2 = nextQ
		cpu2p = nextPlt
	}
}

func use_ioS(ioName string, cpuName string) {
	switch ioName {
	case "1":
		io_cpu(io1, io1p, cpuName)
	case "2":
		io_cpu(io2, io2p, cpuName)
	case "3":
		io_cpu(io3, io3p, cpuName)
	case "4":
		io_cpu(io4, io4p, cpuName)
	default:
		return
	}
}

func io_cpu(io []string, iop []string, cpu string) {
	if cpu == "cpu1" {
		insertQ(io, cpu1, iop, cpu1p)
		cpu1 = ""
		cpu1p = ""
	} else if cpu == "cpu2" {
		insertQ(io, cpu2, iop, cpu2p)
		cpu2 = ""
		cpu2p = ""
	}
	expire(cpu)
}

func use_ioSx(ioName string) {
	fq := ""
	plt := ""
	switch ioName {
	case "1":
		fq, plt = deleteQ(io1, io1p)
	case "2":
		fq, plt = deleteQ(io2, io2p)
	case "3":
		fq, plt = deleteQ(io3, io3p)
	case "4":
		fq, plt = deleteQ(io4, io4p)
	default:
		return
	}
	if fq == "" {
		return
	}

	if cpu1 == "" {
		cpu1 = fq
		cpu1p = plt
		addPriority(plt)
	} else if cpu2 == "" {
		cpu2 = fq
		cpu2p = plt
		addPriority(plt)
	} else {
		if plt == "1" {
			insertQ(ready1, fq, ready1p, plt)
		} else if plt == "2" {
			insertQ(ready2, fq, ready2p, plt)
		} else if plt == "3" {
			insertQ(ready3, fq, ready3p, plt)
		}
	}
}

func main() {
	initx()
	for {
		showProcess()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "new":
			for i := range commandx {
				if i == 0 {
					continue
				}
				//	newProcess(commandx[i])
				if i%2 == 0 {
					newProcess(commandx[i-1], commandx[i])
				}

			}
		case "terminate":
			terminate(commandx[1])
		case "expire":
			expire(commandx[1])
		case "io":
			use_ioS(commandx[1], commandx[2])

		case "iox":
			use_ioSx(commandx[1])

		default:
			fmt.Printf("\nERROR!! PLEASE TRY AGAIN...\n")
		}
	}

}
