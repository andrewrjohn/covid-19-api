## COVID-19 Tracker API
Get global and country-specific information regarding number of cases of COVID-19. Updated every 10 minutes. Data is scraped but responses are dirty cached in memory to reduce response time.

#### Current Endpoints 
- _/api/cases/countries_
- _/api/cases/countries/summary_
- _/api/cases/countries/{country_name}_
  
#### Example Response
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