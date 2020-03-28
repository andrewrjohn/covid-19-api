## COVID-19 Tracker API
Get global and country-specific information regarding number of cases of COVID-19. Updated every 10 minutes. Data is scraped and dirty cached in memory to reduce response time.

### Local Installation
* Clone repo into your Go workspace (i.e. `GOPATH/src`)
* Run `go install`
* If you've added your bin folder to your PATH
  * Run `covid-19-api`
* If you haven't added your bin folder to your PATH
  * Run `GOPATH/bin/covid-19-api`


### Current Endpoints 
- `/api/cases/countries`
- `/api/cases/countries/summary`
- `/api/cases/countries/{country_name}`
  
### Example Response
Note that **new** is always comparative to the previous day.
```
{
    "country_name": "USA",
    "total_cases": 32356,
    "new_cases": 8149,
    "total_deaths": 414,
    "new_deaths": 112,
    "total_recovered": 178,
    "active_cases": 31764,
    "critical_cases": 795,
    "cases_per_million": 98
}
```

#### Resources
- [worldometers.com](https://www.worldometers.info/coronavirus/)