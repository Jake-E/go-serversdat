#go-serversdat
This projects aims to make managing a large servers.dat file easier by providing a server list in tsv format.


##Example Usage
Example `servers.dat` and `server-list.txt` can be found in the [examples](examples) directory.
####List servers:
```
> serversdat -l
┌───┬────────────────────┬────────────────────────┐
│ # │ SERVER NAME        │ IP ADDRESS             │
├───┼────────────────────┼────────────────────────┤
│ 1 │ Local Test         │ localhost              │
│ 2 │ Development Server │ dev.myminecraft.server │
│ 3 │ Minecraft Server   │ play.minecraft.server  │
└───┴────────────────────┴────────────────────────┘
```

####Update servers:
```
> serversdat -u
Updated servers.dat
```

####Export servers:
```
> serversdat -e
```

server-list.txt
```
Local Test	localhost
Development Server	dev.myminecraft.server
Minecraft Server	play.minecraft.server
```

####Arguments:
```
  -e    Alias for -export
  -export
        Exports the current values to a normalized format
  -l    Alias for -list
  -list
        List all servers in table format
  -d string
        Alias for -serverDat (default "./servers.dat")
  -serverDat string
        Path to your Minecraft servers.dat (default "./servers.dat")
  -s string
        Alias for -servers (default "./server-list.txt")
  -servers string
        Path to the servers list (default "./server-list.txt")
  -u    Alias for -update
  -update
        Updates servers.dat file from server list provided in -s

```

##To Do
- Consider using a JSON format for the exported list and list used to update.
- Handle some more potential errors
- Refactor into more manageable functions