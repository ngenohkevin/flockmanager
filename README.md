# Flockmanager

flockmanager is a simple Rest Api that allows consumers to manage
flocks through different stages of production in a farm.
From the number of eggs produced to the number of chicks hatched.

### https://flockmanager.herokuapp.com/v1/kukuchic/

#### kuroiler endpoint [/kuroiler]
#### rainbowrooster endpoint [/rainbowrooster]
#### broiler endpoint [/broiler]
#### layers endpoint [/layers]

### List All flocks [GET]

+ Response 200 (application/json)

        [
            {
              "ID": 6,
              "CreatedAt": "2022-02-06T16:37:25.234829Z",
              "UpdatedAt": "2022-02-21T13:52:44.439305Z",
              "DeletedAt": null,
              "title": "LS1",
              "Production": {
                "eggs": 1000000,
                "dirty": 200,
                "wrong_shape": 5000,
                "weak_shell": 6000,
                "damaged": 20,
                "hatching_eggs": 850000
              },
              "Hatchery": {
                "infertile": 0,
                "early": 0,
                "middle": 0,
                "late": 0,
                "dead_chicks": 0,
                "alive_chicks": 0
              },
              "Premises": {
                "farm": "Kaptagat",
                "house": "45jk"
              }
            }
        ]

### Create a New Flock [POST]

You may create a flock using this action. It takes a JSON
object containing a flock and its data

+ Request (application/json)

        {
            "title": "DF8",
            "Production": {
                "eggs": 10000,
                "dirty": 20,
                "wrong_shape": 10,
                "weak_shell": 20,
                "damaged": 30,
                "hatching_eggs": 9920
            },
            "Hatchery": {
                "infertile": 10,
                "early": 4500,
                "middle": 3000,
                "late": 60,
                "dead_chicks": 50,
                "alive_chicks": 8000
            },
            "Premises": {
                "farm": "Eld",
                "house": "D"
            }
        }

+ Response 201 (application/json)

    + Headers

            Location: /layers/DF8

    + Body

            {
                  "ID": 8,
                  "CreatedAt": "2022-03-23T08:28:58.48333Z",
                  "UpdatedAt": "2022-03-23T08:28:58.48333Z",
                  "DeletedAt": null,
                  "title": "DF8",
                  "Production": {
                    "eggs": 10000,
                    "dirty": 20,
                    "wrong_shape": 10,
                    "weak_shell": 20,
                    "damaged": 30,
                    "hatching_eggs": 9920
                  },
                  "Hatchery": {
                    "infertile": 10,
                    "early": 4500,
                    "middle": 3000,
                    "late": 60,
                    "dead_chicks": 50,
                    "alive_chicks": 8000
                  },
                  "Premises": {
                    "farm": "Eld",
                    "house": "D"
                  }
            }
