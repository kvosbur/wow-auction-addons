## Logic to get buyout for a given item at that snapshot:

- Knowns
1. Auctions can be posted for 3 different durations 12hr, 24hr, 48hr
1. Default auction price is at the current buyout, but you can change this
1. You can cancel an auction at any time, but you still lose the gold you used to post it.
1. The api gives you a snapshot of the auction data on about each hour
1. If an auction appears, it is impossible to know when exactly it was created based on api data
1. If an auction disappears, it is impossible to directly tell if that auction is canceled or sold
1. It is possible to figure out if an auction expired, by tracking when it was added and 

- Logic
1. Can just look at current low price in data to get buyout
1. Extra: Can look for auctions in which the quantities have changes, and use that as a tell that the price hit that point during the iterim

- Unknown
1. How to filter out outliers on the lowend and not use them as the buyout. 
    - Thinking of the kind of auctions that are being used to bait players into selling low.
    - One approach is using IQR: https://www.khanacademy.org/math/statistics-probability/summarizing-quantitative-data/box-whisker-plots/a/identifying-outliers-iqr-rule#:~:text=A%20commonly%20used%20rule%20says,3%20%2B%201.5%20%E2%8B%85%20IQR%20%E2%80%8D%20.
    - Another approach is calculating standard deviation and using zscores: https://www.scribbr.com/statistics/standard-normal-distribution/
    - The first does not take the average into consideration, while the second relies on the data being somewhat normally distributed.


- Requests
1. token: POST https://oauth.battle.net/token with client id/secret using basic auth with form body:: grant_type:client_credentials
2. commodities: GET https://us.api.blizzard.com/data/wow/auctions/commodities?namespace=dynamic-us&locale=en_US with token as bearer token


### Start Database
    - `docker compose -f database/docker-compose.yaml up -d`  psql -d wow-auction-db -U user -W


## TODO
1. Set up recurrence, likely using some docker container
1. Store backup of api response in data file
1. Implement IQR outlier filter
1. Implement z-score outlier filter
