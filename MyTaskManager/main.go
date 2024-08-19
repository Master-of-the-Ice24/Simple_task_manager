package main

import (
	"fmt"
	"strconv"
	//"MyTaskManager/Modules"
	"math"
	"time"
	"github.com/shirou/gopsutil/cpu"
	"os/exec"
	"strings"
	"os"
)

//this function retrieves the usage percentage of the CPU, also known as processor
func CPU_stats() {
	percentage, _ := cpu.Percent(time.Second, false)
	percentageValue :=  float64(percentage[0])
	//percentageString := fmt.Sprintf("%v", percentageValue)
	 percentageString := fmt.Sprintf("%v", math.Floor(float64(percentageValue)*100)/100) + "%"

	fmt.Println("CPU usage: ", percentageString)
}

//this function shows: Total RAM available, free RAM, used RAM and used RAM percentage using the command "cat /proc/meminfo" of GNU/Linux
func virtualMemoryStats() {
	data, _ := os.ReadFile("/proc/meminfo")
	rawText := string(data)
	text := strings.Split(rawText, "\n")
	memorySection := text[:3]

	memoryInfo := []string{}
	for _, element := range memorySection {
		subStrings := strings.Fields(element)
		for i:=0; i<3; i++ {
			memoryInfo = append(memoryInfo, subStrings[i])
		}
	}

	totalMemory, _ := strconv.ParseFloat(memoryInfo[1], 64)
	totalMemory = math.Floor(totalMemory*math.Pow(10, -6)*100)/100

	freeMemory, _ := strconv.ParseFloat(memoryInfo[7], 64)
	freeMemory = math.Floor(freeMemory*math.Pow(10, -6)*10)/10

	usedMemory := math.Floor((totalMemory-freeMemory)*10)/10
	usedMemoryPercentage := math.Floor((usedMemory/totalMemory)*100)
	usedMemoryPercentageString := fmt.Sprintf("%v", usedMemoryPercentage) + "%"

	fmt.Println("Total memory: ", totalMemory, "GB")
	fmt.Println("Free memory: ", freeMemory, "GB")
	fmt.Println("Used memory: ", usedMemory, "GB")
	fmt.Println("Used memory percentage: ", usedMemoryPercentageString)
}

//this function merges CPU_stats() and virtualMemoryStats()
func resourcesStats() {
	fmt.Println("\n\n--------------------SYSTEM RESOURSES USAGE--------------------")
	CPU_stats()
	virtualMemoryStats()
	fmt.Println("--------------------------------------------------------------")
}

//this function finds CPU information using the command "lscpu" of GNU/Linux
func CPU_info() {
	command := exec.Command("lscpu")
	info, _ := command.Output()

	fmt.Println("--------------------CPU INFORMATION--------------------")
	fmt.Println(string(info))
	fmt.Println("-------------------------------------------------------")
}

func main() {
	fmt.Println("Welcome on Master_of_the_ice24's Task Manager\n\n")

	CPU_info()	
	resourcesStats()	
}