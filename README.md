# wslroutesvc

Simple service to remove routes in the Windows routing table when they conflict with existing WSL virtual interface routes.

## Usage
Install the service (as Administrator):
```powershell
.\wslroutesvc.exe install
```

After installing the service, start it and set it to automatic within services.msc or enable/start it on the command line.
