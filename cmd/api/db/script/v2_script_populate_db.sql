with customer_json (doc) as (
   values 
    ('[
  {
    "id": 1,
    "fullname": "Renan Machado",
    "email": "renanoliveira2199@gmail.com",
    "password": "lot@199Z",
    "socialname": "Renan"
  },
  {
    "id": 2,
    "fullname": "Hannah Moura",
    "email": "hannahmoura@gmail.com",
    "password": "hannahmoura",
    "socialname": "Hannah"
  }
]'::json)
)
insert into customer (id, fullname,email,password,socialname)
select p.*
from customer_json l
cross join lateral json_populate_recordset(null::customer, doc) as p

----

with customer_json (doc) as (
   values 
    ('[
  {
    "id": 1,
    "user_id": 1,
    "title": "Salário",
    "description": "Salario",
    "type": "Trabalho",
    "value": 6000.99,
    "datetime": "2024-05-31T00:00:00Z",
    "createdAt": "2024-05-01T15:36:39Z"
  },
  {
    "id": 2,
    "user_id": 2,
    "title": "Salário",
    "description": "Salario",
    "type": "Trabalho",
    "value": 12599.01,
    "datetime": "2024-05-31T00:00:00Z",
    "createdAt": "2024-05-01T15:36:39Z"
  }
]'::json)
)
insert into incomes (id,	user_id,	title,	description,	type,	value,	datetime)
select p.*
from customer_json l
cross join lateral json_populate_recordset(null::incomes, doc) as p

----

with customer_json (doc) as (
   values 
    ('[
  {
    "id": 1,
    "user_id": 1,
    "title": "Netflix",
    "description": "Assinatura netflix mensal",
    "type": "entreterimento",
    "value": 40.15,
    "datetime": "2024-05-31 13:04:05"
  },
  {
    "id": 2,
    "user_id": 2,
    "title": "Gympass",
    "description": "Assinatura Gympass mensal",
    "type": "Saúde",
    "value": 180.89,
    "datetime": "2024-05-3 00:00:00"
  },
  {
    "id": 3,
    "user_id": 1,
    "title": "Gasolina",
    "description": "gasolina pro focus",
    "type": "locomoção",
    "value": 212.48,
    "datetime": "2024-05-31 13:04:05"
  },
  {
    "id": 4,
    "user_id": 2,
    "title": "Mercado",
    "description": "Compras semana 1 de junho",
    "type": "mercado",
    "value": 97.67,
    "datetime": "2024-05-31 13:04:05"
  }
]
'::json)
)
insert into expenses (id,	user_id,	title,	description,	type,	value,	datetime)
select p.*
from customer_json l
cross join lateral json_populate_recordset(null::expenses, doc) as p