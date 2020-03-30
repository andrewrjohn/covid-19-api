## COVID-19 Tracker API
Get global and country-specific information regarding number of cases of COVID-19. Updated every 10 minutes. Data is scraped and dirty cached in memory to reduce response time.

### Local Installation
* Clone repo into your Go workspace (i.e. `GOPATH/src`)
* Run `go install`
* If you've added your bin folder to your PATH
  * Run `covid-19-api`
* If you haven't added your bin folder to your PATH
  * Run `GOPATH/bin/covid-19-api`

### Hot-Reloading for Development (optional)
To increase your productivity and take advantage of server hot-reloading, do the following steps:
* Download the utility: `go get github.com/codegangsta/gin`
* Start the dev server one of two ways:
  * Use the included .sh file by running `./run.sh` (you'll probably have to add permissions)
  * Run `gin -a 8080 run main.go`
* Hot-Loading Server will be running on port 3000



### Current Endpoints 
- `/api/cases/countries`
- `/api/cases/countries/summary`
- `/api/cases/countries/{country_name}`
  
### Example Response
Note that **new** is always comparative to the previous day.
```
{
  "country_name": "USA",
  "total_cases": 145542,
  "active_cases": 138347,
  "new_cases": 2051,
  "critical_cases": 2972,
  "total_deaths": 2616,
  "new_deaths": 33,
  "death_rate": "1.80%",
  "total_recovered": 4579,
  "cases_per_million": 440
}
```

#### Resources
- [worldometers.com](https://www.worldometers.info/coronavirus/)