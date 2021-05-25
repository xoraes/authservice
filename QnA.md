How long did this assignment take?
> Due to the disk failure on my laptop, it took me close to 8 hours as I had to rewrite much of the application

What was the hardest part?
> The part that took longest was to set up postgres and the auth service to work together with docker-compose

Did you learn anything new?
> Fortunately, I have done most of this before. However, I had not used postgres db in the past. 

Is there anything you would have liked to implement but didn't have the time to?
> * Better input validation. This will also prevent potential sql injection attack. 
> * Improving testability by making better use of interfaces for IoC/mocks
> * Writing tests and deployment code (ci/cd)
> * Application caching for performance
> * Allowing switching to nossl and debug mode
> * Refresh token and delete user features

What are the security holes (if any) in your system? If there are any, how would you fix them?
> * We care about security both in transit and at rest. 
> * Data in DB should be encrypted and data should be secured while in transit between db and app (using ssl for example).  
> * Replay attack/Man in the middle could change first and last name if a bad actor got hold of the token. We can use refresh tokens and lower the token ttl. The token could also be generated using the users ip-address / geo to reduce the attack surface.
> * While I added TLS, security credentials should be in a vault and only accessed during deployment. I added the ssl and jwt keys to git repo for eval purposes, however there are better ways to handle that.
> * .env should not contain sensitive information. SSL private key and JWT key should be protected in a vault.
> * SQL Inject Attack. We can prevent it with better input validation. I already use paramterized queries.  

Do you feel that your skills were well tested? 
> * Yes
     