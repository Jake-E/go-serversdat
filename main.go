package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/seebs/nbt"
	"io/ioutil"
	"os"
)

// Server Struct
type Server struct {
	name string
	ip   string
}

func main() {
	serverDat := flag.String("serverDat", "./servers.dat", "Path to your Minecraft servers.dat")
	listPtr := flag.String("servers", "./server-list.txt", "Path to the servers list")
	export := flag.Bool("export", false, "Exports the current values to a normalized format")
	list := flag.Bool("list", false, "List all servers in table format")
	update := flag.Bool("update", false, "Updates servers.dat file from server list provided in -s")

	flag.StringVar(serverDat, "d", *serverDat, "Alias for -serverDat")
	flag.StringVar(listPtr, "s", *listPtr, "Alias for -servers")
	flag.BoolVar(export, "e", *export, "Alias for -export")
	flag.BoolVar(list, "l", *list, "Alias for -list")
	flag.BoolVar(update, "u", *update, "Alias for -update")
	flag.Parse()


	if *update {
		list, _ := ioutil.ReadFile(*listPtr)
		s := bytes.NewReader(list)

		reader := csv.NewReader(s)
		reader.Comma = '\t'

		reader.FieldsPerRecord = -1

		serverList, ok := reader.ReadAll()
		if ok != nil {
			fmt.Println(ok)
			os.Exit(1)
		}

		var servers []Server
		var tmp Server

		for _, each := range serverList {
			tmp.name = each[0]
			tmp.ip = each[1]
			servers = append(servers, tmp)
		}

		var srvs []nbt.Compound
		for _, each := range servers {
			c := make(nbt.Compound)
			c["ip"] = nbt.String(each.ip)
			c["name"] = nbt.String(each.name)
			srvs = append(srvs, c)
		}

		addServers := nbt.MakeCompoundList(srvs)

		sC := make(nbt.Compound)
		sC["servers"] = addServers

		fileW, ok := os.OpenFile(*serverDat, os.O_RDONLY|os.O_CREATE, 0666)
		if ok != nil {
			fmt.Println(ok)
			os.Exit(1)
		}
		nbt.StoreUncompressed(fileW, sC, "")
		fmt.Println("Updated servers.dat")
		os.Exit(0)
	}

	if *list || *export {
		file, _ := ioutil.ReadFile(*serverDat)
		z := bytes.NewReader(file)

		data, _, err := nbt.LoadUncompressed(z)
		if err != nil {
			fmt.Println("Failed to read servers.dat at path " + string(file))
			fmt.Print(err)
			os.Exit(1)
		}

		n := data.(nbt.Compound)
		serverData, ok := nbt.GetList(n["servers"])

		if !ok {
			fmt.Println("Failed to decode servers.dat at path " + string(file))
			os.Exit(1)
		}

		servers, ok := serverData.GetCompoundList()
		if !ok {
			fmt.Println("Failed decoding server compound list")
			os.Exit(1)
		}

		var t table.Writer
		if *list {
			fmt.Print("Found ")
			fmt.Print(serverData.Length())
			fmt.Println(" server(s):")

			t = table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"#", "Server Name", "IP Address"})
		}
		var o string
		for i := 0; i < serverData.Length(); i++ {
			server, ok := nbt.GetCompound(servers[i])
			if !ok {
				fmt.Println("Error in servers.dat entry")
				os.Exit(1)
			}

			ip, _ := nbt.GetString(server["ip"])
			name, _ := nbt.GetString(server["name"])

			if *list {
				t.AppendRow([]interface{}{i + 1, name, ip})
			}
			if *export {
				o = o + string(name) + "\t" + string(ip)
				if i+1 != serverData.Length() {
					o = o + "\n"
				}
			}
		}
		if *list {
			t.SetStyle(table.StyleLight)
			t.RenderMarkdown()
		}
		if *export {
			ok := ioutil.WriteFile(*listPtr, []byte(o), 0644)
			if ok != nil {
				fmt.Println(ok)
				os.Exit(1)
			}
			fmt.Print("Exported ")
			fmt.Print(serverData.Length())
			fmt.Println(" server(s)")
		}
		os.Exit(0)
	}

	flag.PrintDefaults()
}