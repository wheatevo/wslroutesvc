# WSL Routing Conflict Service (wslroutesvc.exe)

Simple service to remove routes in the Windows routing table when they conflict with existing WSL virtual interface routes.

## Usage
Install the service (as Administrator):
```powershell
.\wslroutesvc.exe install
```

This command will start the service and set it to automatic startup.
