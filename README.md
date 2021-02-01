# microservices

<a href="https://www.buymeacoffee.com/mattvogel" target="_blank">
	<img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" >
</a>

## Auth Service
### Description
User management service, create users, serve JWT and store existing sessions in Redis
### Routes
 * Create User - /v1/user - POST
 * Login - /v1/user/login - POST
 * Logout - /v1/user/logout - GET
 * Token Refresh - /v1/user/refresh - POST

## Device Service
### Description
Device configuration service, management of new and old devices
### Routes
 * Add Device - /v1/device - POST
 * Get Device By ID - /v1/device/:deviceID - GET
 * Get Device by Owner - /v1/device - GET
 * Update Device - /v1/device/:deviceID - PUT

## Condition service
### Description
Collect and return environmental conditions per device.
### Routes
 * Send Temperature - /v1/device/:deviceID/temp - POST
 * Get Conditions - /v1/device/:deviceID/conditions - GET

## Database setup
### Timescale
The database that is currently in use is the Timescale PG12 docker container. This can be easily swapped out for a regular install of Timescale if needed.

### Redis
Redis is currently just a secured instance container without any other considerations.
Typical use is that we generate tokens and place them in redis with a TTL, intention  for this is to allow administrative users to see how many users are logged into the service at any specific time.
