package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID     int    `json:"ID"`
	Input  string `json:"Task"`
	Status bool   `json:"Status"`
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}
}

var job = []Task{}

func add(task string) {
	nextID := 1
	for _, t := range job {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}
	job = append(job, Task{ID: nextID, Input: task, Status: false})
}
func load() {
	data, err := os.ReadFile("tasks.json")
	if os.IsNotExist(err) {
		return
	}
	check(err)

	err = json.Unmarshal(data, &job)
	check(err)
}
func save() {
	data, err := json.MarshalIndent(job, "", "    ")
	check(err)
	err = os.WriteFile("tasks.json", data, 0644)
	check(err)
	fmt.Println("Tugas telah disimpan ke tasks.json")
}
func remove(id int) bool {
	for i, t := range job {
		if t.ID == id {
			job = append(job[:i], job[i+1:]...)
			return true
		}
	}
	return false
}

// func update(id int)
func update(id int) bool {
	for i := range job {
		if job[i].ID == id {
			job[i].Status = !job[i].Status
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Perintah tidak dikenali")
		return
	}
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Tugas tidak ditemukan")
			return
		}
		if len(os.Args) > 3 {
			fmt.Println("Error: Argumen tugas terlalu banyak!")
			fmt.Println("Gunakan tanda kutip jika tugas lebih dari 1 kata. Contoh: add \"makan malem bersama\"")
			return
		}
		load()
		add(os.Args[2])
		save()
		fmt.Println("Tugas telah diterima dengan ID", job[len(job)-1].ID)
	case "list":
		fmt.Println("Daftar tugas :")
		load()
		for _, t := range job {
			status := " "
			if t.Status {
				status = "x"
			}
			fmt.Printf("%d. [%s] %s\n", t.ID, status, t.Input)
		}
	case "update":
		if len(os.Args) != 3 {
			fmt.Println("Gunakan: update <ID>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID harus berupa angka")
			return
		}
		load()
		if !update(id) {
			fmt.Println("Tugas dengan ID", id, "Tidak Ditemukan")
			return
		}
		save()
		fmt.Println("Tugas Telah diperbarui")
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Gunakan: Delete <ID>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID harus berupa angka")
			return
		}
		load()
		if !remove(id) {
			fmt.Println("Tugas dengan ID", id, "Tidak Ditemukan")
			return
		}
		save()
		fmt.Println("Tugas Telah Dihapus")
	default:
		fmt.Println("Hanya ada add, list, update, delete as command")
	}

}
