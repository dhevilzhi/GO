package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func find(key string, personDB map[string]PersonInfo) (info PersonInfo, ok bool) {
	info, ok = personDB[key]
	return
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	//	personDB = make(map[string] PersonInfo, 100) // 100为map的存储能力

	// 插入数据
	personDB["1"] = PersonInfo{"1", "Tom", "Room 1"}
	personDB["2"] = PersonInfo{"2", "Jack", "Room 2"}

	/*
		// 创建并初始化map
		personDB = map[string]PersonInfo{
			"1": PersonInfo{"1", "Tom", "Room 1"},
			"2": PersonInfo{"2", "Jack", "Room 2"}
		}
	*/

	// 查找数据
	person, ok := find("1", personDB)
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1, Address is", person.Address)
	} else {
		fmt.Println("Did not find person with ID 1.")
	}

	// 删除数据
	delete(personDB, "1")
	person, ok = find("1", personDB)
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1, Address is", person.Address)
	} else {
		fmt.Println("Did not find person with ID 1.")
	}
}
