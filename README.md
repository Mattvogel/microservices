# microservices

<a href="https://www.buymeacoffee.com/mattvogel" target="_blank">
	<img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" >
</a>

## Auth Service
###  Routes
 * Create User - /v1/user - POST
 * Login - /v1/user/login - POST
 * Logout - /v1/user/logout - GET
 * Token Refresh - /v1/user/refresh - POST

## Device Service
 * Add Device - /v1/device - POST
 * Get Device By ID - /v1/device/:deviceID - GET
 * Get Device by Owner - /v1/device - GET
 * Update Device - /v1/device/:deviceID - PUT

## Condition service
 * Send Temperature - /v1/device/:deviceID/temp - POST
 * Get Conditions - /v1/device/:deviceID/conditions - GET

## Database setup
The database that is currently in use is the Timescale PG12 docker container. This can be easily swapped out for a regular install of Timescale if needed.

