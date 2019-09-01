# Salamantex test back-end

### Stack:

- **Go** Base programming language
- **Postgresql** Database to store data
- **Redis** Basic queue management (POP/PUSH)
- **Docker** Make the back-end work in containers
- **Docker Compose** Connect all services together

### Running the app

Simply start the app with the docker-compose:

```shell script
docker-compose up -d
```

### Database structure
![users](https://user-images.githubusercontent.com/16473040/64076211-2c03eb00-ccd7-11e9-90dd-9a5ffd64a659.png)
![tx](https://user-images.githubusercontent.com/16473040/64076217-3b833400-ccd7-11e9-803d-4d74962a20c7.png)

### Endpoints

#### Register

It will give you your JWT token.

**URL**: /api/v1/user/register

**Method**: POST

**Needs auth?** No

**Header**: `none`

**Request Body**:

```json
{
  "name": "NAME",
  "email": "name@mail.com",
  "password": "password"
}
```


#### Login

It will give you your JWT token.

**URL**: /api/v1/user/login

**Method**: POST

**Needs auth?** No

**Header**: `none`

**Request Body**:

```json
{
  "email": "name@mail.com",
  "password": "password"
}
```


#### Add Currency

By adding an address, It will give you 10 units of that currency for testing.

**URL**: /api/v1/currency/add

**Method**: POST

**Needs auth?** Yes

**Header**:

```json
{
  "Authorization": "Bearer TOKEN"
}
```

**Request Body**:

```json
{
  "type": 1,
  "address": "1F1tAaz5x1HUXrCNLbtMDqcw6o5GNn4xqX"
}
```

*Note*:

Type `1` -> Bitcoin

Type `2` -> Ethereum


#### Send Transaction

**URL**: /api/v1/tx/send

**Method**: POST

**Needs auth?** Yes

**Header**:

```json
{
  "Authorization": "Bearer TOKEN"
}
```

**Request Body**:

```json
{
  "type": 1,
  "amount": 5,
  "to": "friend@mail.com"
}
```

*Note*: 

Type `1` -> Bitcoin

Type `2` -> Ethereum


#### Transactions History

**URL**: /api/v1/tx/history

**Method**: GET

**Needs auth?** Yes

**Header**:

```json
{
  "Authorization": "Bearer TOKEN"
}
```

**Request Body**: `none`


#### Transaction state

**URL**: /api/v1/tx/state/`<ID>`

**Method**: GET

**Needs auth?** Yes

**Header**:

```json
{
  "Authorization": "Bearer TOKEN"
}
```

**Request Body**: `none`