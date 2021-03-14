# WSL Routing Conflict Service (wslroutesvc.exe)

Simple service to remove routes in the Windows routing table when they conflict with existing WSL virtual interface routes.

## Usage
* Download the service:
[wslroutesvc.exe](https://github.com/wheatevo/wslroutesvc/releases/latest/download/wslroutesvc.exe)

* Install the service (as Administrator):
```powershell
.\wslroutesvc.exe install
```
> This command will start the service and set it to automatic startup.
