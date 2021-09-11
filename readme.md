This is rest api CRUD service.
The service has 4 APIs, first api is CreateUser api.
It accepts the request:
{
    "first_name": "Joe",
    "last_name": "Parker",
    "age": 29,
    "mail":"joe@gmail.com"
}
and returns response:
{
    "msg":"user created"
}
Second api is UpgradeUserByUUID accepts:
{
    "id": 1,
    "first_name": "Joe",
    "last_name": "Parker",
    "age": 29,
    "mail":"joe@gmail.com"
}
response is:
{
    "msg": "user is upgraded"
}
The last api is GetUserByUUID, request object of this api is:
{
    "id": 1
}
and the response is:
{
    "id": 1,
    "first_name": "Joe",
    "last_name": "Parker",
    "age": 29,
    "mail":"joe@gmail.com"
}