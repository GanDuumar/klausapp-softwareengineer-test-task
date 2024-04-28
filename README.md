# Software Engineer Test Task

As a test task for [Klaus](https://www.klausapp.com) software engineering position we ask our candidates to build a small [gRPC](https://grpc.io) service using language of their choice. Preferred language for new services in Klaus is [Go](https://golang.org).

The service should be using provided sample data from SQLite database (`database.db`).
* Database in internal/db/database.db

Please fork this repository and share the link to your solution with us.

### Tasks

1. Come up with ticket score algorithm that accounts for rating category weights (available in `rating_categories` table). Ratings are given in a scale of 0 to 5. Score should be representable in percentages from 0 to 100. 
1.1 Solved it as a db.go function, due to DB access request.
1.2 Could've perhaps cached the ratings, when using multiple times.

2. Build a service that can be queried using [gRPC](https://grpc.io/docs/tutorials/basic/go/) calls and can answer following questions:
2.1 Service up, using gRPC and I myself used Postman for requests.
2.2 As it was not defined, I presume the period is in "Y-m-d" format and no time restriction required (From 00:00:00 to 23:59:59 all valid)

    * **Aggregated category scores over a period of time**
    
        E.g. what have the daily ticket scores been for a past week or what were the scores between 1st and 31st of January.
        For periods longer than one month weekly aggregates should be returned instead of daily values.
        From the response the following UI representation should be possible:

        | Category | Ratings | Date 1 | Date 2 | ... | Score |
        |----|----|----|----|----|----|
        | Tone | 1 | 30% | N/A | N/A | X% |
        | Grammar | 2 | N/A | 90% | 100% | X% |
        | Random | 6 | 12% | 10% | 10% | X% |
        
        * Note: I presume the example "Tone" is just as a structure example, as in the database.db I have: Spelling, Grammar, GDPR, Randomness 
        *  Endpoint: GetRatingsBetweenDatesAggregated
        *  Body example: 
        ```
        {
            "start_date": "2019-07-05",
            "end_date": "2019-07-17"
        }
        ```
        * Response example:
        ```
        {
        "rating_group": [
            {
                "ratings": [
                    {
                        "date_or_week": "2019-07-06",
                        "rating": 55
                    },
                    ...
                    {
                        "date_or_week": "2019-07-14",
                        "rating": 73
                    }
                ],
                "id": 3,
                "name": "GDPR",
                "count": 12
            },
            ...
        ```
     
    * **Scores by ticket**

        Aggregate scores for categories within defined period by ticket.

        E.g. what aggregate category scores tickets have within defined rating time range have.

        | Ticket ID | Category 1 | Category 2 |
        |----|----|----|
        | 1   |  100%  |  30%  |
        | 2   |  30%  |  80%  |
        
        *  Endpoint: GetCategoryRatingsBetweenDatesAggregated
        *  Body example: 
        ```
        {
            "start_date": "2019-07-05",
            "end_date": "2019-07-17"
        }
        ```
        * Response example:
        ```
        {
        "rating_group": [
        {
            "ratings": [
                {
                    "date_or_week": "2019-34",
                    "rating": 47
                },
                {
                    "date_or_week": "2019-26",
                    "rating": 54
                },
                ...
            ],
            "id": 1,
            "name": "Spelling",
            "count": 12
        },
        ...
        ```

    * **Overall quality score**

        What is the overall aggregate score for a period.

        E.g. the overall score over past week has been 96%.
        
        * Endpoint: GetOverallRatingsBetweenDates
        * Body example:
        ```
        {
            "start_date": "2019-07-05",
            "end_date": "2019-07-17"
        }
        ```
        * Response example:
        ```
        {
            "overall_rating": 35
        }
        ```

    * **Period over Period score change**

        What has been the change from selected period over previous period.

        E.g. current week vs. previous week or December vs. January change in percentages.

        * NB: I presume that I can provide start-end dates and it tells me the changes between those. Other choice would be to aggregate months and weeks and diff them, giving only input for example start "2024-01". My solution seemed logical to me at the time, giving start date and end date, tries to diff between them in general.
        * Endpoint: GetOverallRatingsChangeBetweenDates
        * Body example:
        ```
        {
            "start_date": "2019-07-05",
            "end_date": "2019-07-17"
        }
        ```
        * Response example:
        ```
        {
            "percentage_change": 35
        }
        ```

### Bonus

1.  How would you build and deploy the solution?
    * At Klaus we make heavy use of containers and [Kubernetes](https://kubernetes.io).
    * I would use Docker container, as I'm most used to it. Maybe setup a CI/CD to AWS Beanstalk. Automate the pipeline with workers. Auto-deploy. Though that would be only doable if down-time is allowed (During the Beanstalk image deploy). For better constant uptime, would think of other methods. Maybe for simple project even a simple git hook system (But would not preffer it myself). Would need more research for constant Go deployment in AWS.
    
### Technical notes

* Builds and scripts
    * There is a "build.sh" file for proto builds
    * Running "**go build -o server.exe ./cmd/server***" in root path creates the "server.exe" file that needs to be run for the program to work
    * Code was developed using Visual Studio Code (1.88.1) in Windows 11
    * Postman was used locally for requests, importing the proto file for the structure

* Notes on the code by the developer
    * Things I would improve, if possible
        * Make unit tests, to make sure the data is indeed correct (Manual check currently implemented) 
        * gRPC response bodies sorted more logically (id, name and others before longer lists/objects)
        * Re-use methods and avoid calculations (for example gRPC handlers)
        * Maybe alter the project structure and naming to better reflect go standards
        * Add more automated scripts for the whole project (Proto files, go build and more)
        * Some code clean-up could be done, in terms of excess comments or methods in general
