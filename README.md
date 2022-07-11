# gcp-rest-api
Application that interacts with GCP PostgreSQL and Firestore.

The application is implemented in order to consolidate knowledge about the communication of the application deployed on CloudRun and GCP databases, as well as the implementation of a Clean Architecture.

`/users*` enpoints interact to the GCP Firestore.

`/authors*` enpoints interact to the PostgreSQL.

### REST API
Endpoint | Method | Response codes | Description
--- | --- | --- | ---
*/users* | GET | 200, 500 | Get list of users
*/users/:id* | GET | 200, 404, 500 | Get user by id
*/users* | POST | 201, 422, 500 | Create user
*/users/:id* | PUT | 204, 404, 400, 422, 500 | Update user
*/users/:id* | DELETE | 204, 404, 500 | Delete user by id
 |  |  | 
*/authors* | GET | 200, 500 | Get list of authors
*/authors/:id* | GET | 200, 404, 500 | Get author by id
*/authors* | POST | 201, 422, 500 | Create author
*/authors/:id* | PUT | 204, 404, 400, 422, 500 | Update author
*/authors/:id* | DELETE | 204, 404, 500 | Delete author by id