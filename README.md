# Oficina de código: Desafio Go II

### Objetivo:
A seguir, apresenta-se um desafio integrador que nos permitirá avaliar todos os tópicos abordados no curso.

## Sistema de marcação de consultas
Pretende-se implementar uma API que permita administrar a marcação de consultas para uma clínica odontológica, que deve atender aos requerimentos a seguir:

### Administração de dados de dentistas:
  Listar, adicionar, alterar ou excluir dentitas. Registrar seus **sobrenome, nome e matrícula**. Solicita-se o desenvolvimento de um **CRUD** para a entidade **Dentista**.
  #### Métodos
  Requisições para a API devem seguir os padrões:

  | Método   | Descrição                                                    |
  |----------|--------------------------------------------------------------|
  | `POST`   | Adicionar dentista.                                          |
  | `GET`    | Trazer dentista pelo seu ID.                                 |
  | `PUT`    | Atualizar dentista.                                          |
  | `PATCH`  | Atualizar um dentista através de algum dos seus campos.      |
  | `DELETE` | Excluir dentista.                                            |
 ### Request / Responses
 ### Adicionar dentista [POST /api/v1/dentists]
  + Request (application/json)
      + Header

              {"SECRET_TOKEN": "TOKEN"}
      + Body

              {
                "surname": "Silva",
                "name": "Alguem",
                "license_number": "123456"
              }
  + Response 201 (application/json)
      + Body

             {
                 "data": {
                     "id": 1,
                     "surname": "Silva",
                     "name": "Alguem",
                     "license_number": "123456"
                 }
             }
  + Response 400 (application/json)
      + Body

             {
                 "code": 400,
                 "status": "error",
                 "message": "invalid data provided",
                 "time_stamp": "2022-12-08 18:42:45.645419 -0300 -03 m=+3.702563901"
             }
  + Response 401 (application/json)
      + Body

             {
                 "code": 401,
                 "status": "error",
                 "message": "Token not found",
                 "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
             }
 ### Listar dentistas [GET /api/v1/dentists]
  + Request (application/json)
    + Header

            {"SECRET_TOKEN": "TOKEN"}
  + Response 200 (application/json)
      + Body

             {
                "data": [
                    {
                        "id": 1,
                        "surname": "Silva",
                        "name": "Alguem",
                        "license_number": "123456"
                    },
                    {
                        "id": 2,
                        "surname": "Silva",
                        "name": "João",
                        "license_number": "00000001"
                    },
                    {
                        "id": 3,
                        "surname": "Arantes",
                        "name": "Guilherme",
                        "license_number": "000030001"
                   }
                ]
             }
  + Response 400 (application/json)
      + Body

             {
                 "code": 400,
                 "status": "error",
                 "message": "invalid request",
                 "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
             }
  + Response 401 (application/json)
      + Body

             {
                 "code": 401,
                 "status": "error",
                 "message": "Token not found",
                 "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
             }
 ### Buscar dentista [GET /api/v1/dentists/{id}]
  + Request (application/json)
      + Header

              {"SECRET_TOKEN": "TOKEN"}
  + Response 200 (application/json)
     + Body

            {
                "data": {
                    "id": 1,
                    "surname": "Silva",
                    "name": "Alguém",
                    "identity_number": "123456"
                }
            }
  + Response 400 (application/json)
     + Body

            {
                "code": 400,
                "status": "error",
                "message": "invalid id provided",
                "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
            }
  + Response 401 (application/json)
     + Body

            {
                "code": 401,
                "status": "error",
                "message": "Token not found",
                "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
            }
  + Response 404 (application/json)
     + Body

            {
                "code": 404,
                "status": "error",
                "message": "dentist not found",
                "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
            }
 ### Atualizar dentista [PUT /api/v1/dentists/{id}]
 + Request (application/json)
     + Header

             {"SECRET_TOKEN": "TOKEN"}  
    + Body

            {
              "surname": "Almeida Silva",
              "name": "João",
              "license_number": "05000001"
            }
 + Response 200 (application/json)
     + Body

            {
                "id": 1,
                "surname": "Almeida Silva",
                "name": "Alguem",
                "license_number": "05000001"
            }
 + Response 400 (application/json)
     + Body

            {
                "code": 400,
                "status": "error",
                "message": "invalid id provided",
                "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
            }
 + Response 401 (application/json)
     + Body

            {
                "code": 400,
                "status": "error",
                "message": "Token not found",
                "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
            }
 + Response 404 (application/json)
     + Body

            {
                "code": 404,
                "status": "error",
                "message": "dentist not found",
                "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
            }
 ### Atualizar dentista [PATCH /api/v1/dentists/{id}]
 + Request (application/json)
   + Header

           {"SECRET_TOKEN": "TOKEN"}
   + Body

           {
              "surname": "Silva",
           }
 + Response 200 (application/json)
     + Body

            {
                "id": 1,
                "surname": "Silva",
                "name": "João",
                "license_number": "05000001"
            }
 + Response 400 (application/json)
     + Body

            {
                "code": 400,
                "status": "error",
                "message": "invalid id provided",
                "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
            }
 + Response 400 (application/json)
     + Body

            {
                "code": 400,
                "status": "error",
                "message": "invalid data provided",
                "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
            }
 + Response 401 (application/json)
     + Body

            {
                "code": 400,
                "status": "error",
                "message": "Token not found",
                "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
            }
 + Response 404 (application/json)
     + Body

            {
                "code": 404,
                "status": "error",
                "message": "dentist not found",
                "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
            }
 ### Excluir dentista [DELETE /api/v1/dentists/{id}]
   + Request (application/json)
      + Header

            {"SECRET_TOKEN": "TOKEN"}
   + Response 200 (application/json)
       + Body

              {
                  "success": "dentist deleted"
              }
   + Response 400 (application/json)
       + Body

              {
                  "code": 400,
                  "status": "error",
                  "message": "invalid id provided",
                  "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
              }
   + Response 401 (application/json)
       + Body

             {
               "code": 400,
               "status": "error",
               "message": "Token not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
             }
   + Response 404 (application/json)
       + Body

              {
                  "code": 404,
                  "status": "error",
                  "message": "unable to delete, dentist not found",
                  "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
              }

### Administração de dados de pacientes:
  Listar, adicionar, alterar ou excluir pacientes. Registrar seus **sobrenome, RG, nome e matrícula**. Solicita-se o desenvolvimento de um **CRUD** para a entidade **Paciente**.
  #### Métodos
  Requisições para a API devem seguir os padrões:

  | Método   | Descrição                                               |
  |----------------------------------------------------------|---------|
  | `POST`   | Adicionar paciente.                                     |
  | `GET`    | Trazer paciente pelo seu ID.                            |
  | `PUT`    | Atualizar paciente.                                     |
  | `PATCH`  | Atualizar um paciente através de algum dos seus campos. |
  | `DELETE` | Excluir paciente.                                       |
 ### Request / Responses
  ### Adicionar paciente [POST /api/v1/patients]
  + Request (application/json)
    + Header

            {"SECRET_TOKEN": "TOKEN"}
    + Body
    
            {
              "surname": "Rocha",
              "name": "Julia",
              "identity_number": "00000001",
              "created_at": "07/12/2022 23:00",    
            }
  + Response 201 (application/json)
    + Body

           {
              "id": 1,
              "surname": "Rocha",
              "name": "Julia",
              "identity_number": "00000001",
              "created_at": "07/12/2022 23:00",  
           }
  + Response 400 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "invalid data provided",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
  + Response 401 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "invalid data provided",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901""
           }
  ### Buscar paciente [GET /api/v1/patients/{id}]
  + Request (application/json)
    + Header

            {"SECRET_TOKEN": "TOKEN"}
  + Response 200 (application/json)
    + Body

           {
                 "data": {
                     "id": 1,
                     "surname": "Queiroz",
                     "name": "Julia",
                     "identity_number": "00000001",
                     "created_at": "07/12/2022 11:00"
                 }
            }
  + Response 400 (application/json)
    + Body

          {
                "code": 400,
                "status": "error",
                "message": "invalid id provided",
                "time_stamp": "2022-12-07 14:55:13"
          }
  + Response 401 (application/json)
    + Body

          {
              "code": 400,
              "status": "error",
              "message": "Token not found",
              "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901""
          }
  + Response 404 (application/json)
    + Body

          {
               "code": 404,
               "status": "error",
               "message": "patient not found",
               "time_stamp": "2022-12-07 14:55:13"
          }
  ### Atualizar paciente [PUT /api/v1/patients/{id}]
  + Request (application/json)
    + Header

            {"SECRET_TOKEN": "TOKEN"}
    + Body

             {
               "surname": "Queiroz",
               "name": "Julia",
               "identity_number": "00000001",
               "created_at": "07/12/2022 11:00"  
            }
  + Response 200 (application/json)
    + Body

            {
                 "data": {
                     "id": 1,
                     "surname": "Queiroz",
                     "name": "Julia",
                     "identity_number": "00000001",
                     "created_at": "07/12/2022 11:00"
                 }
            }
  + Response 400 (application/json)
    + Body

            {
                "code": 400,
                "status": "error",
                "message": "invalid id provided",
                "time_stamp": "2022-12-07 14:55:13"
            }
  + Response 400 (application/json)
    + Body

            {
                "code": 400,
                "status": "error",
                "message": "invalid data provided",
                "time_stamp": "2022-12-07 14:55:13"
            }
    + Response 401 (application/json)
    + Body

            {
                "code": 401,
                "status": "error",
                "message": "Token not found",
                "time_stamp": "2022-12-07 14:55:13"
            }
  + Response 404 (application/json)
    + Body

            {
                "code": 404,
                "status": "error",
                "message": "patient not found",
                "time_stamp": "2022-12-07 14:55:13"
            }
  ### Atualizar paciente [PATCH /api/v1/patients/{id}]
  + Request (application/json)
    + Header
       
            {"SECRET_TOKEN":"TOKEN"}
    + Body

             {
               "identity_number": "04333330",
             }
  + Response 200 (application/json)
    + Body

             {
                 "data": {
                     "id": 1,
                     "surname": "Queiroz",
                     "name": "Julia",
                     "identity_number": "04333330",
                     "created_at": "07/12/2022 11:00"
                 }
             }
  + Response 400 (application/json)
    + Body

               {
                   "code": 400,
                   "status": "error",
                   "message": "invalid id provided",
                   "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
               }
    + Response 400 (application/json)
    + Body

               {
                   "code": 400,
                   "status": "error",
                   "message": "invalid data provided",
                   "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
               }
  + Response 401 (application/json)
     + Body

               {
                   "code": 401,
                   "status": "error",
                   "message": "Token not found",
                   "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
               }
  + Response 404 (application/json)
    + Body

               {
                   "code": 404,
                   "status": "error",
                   "message": "patient not found",
                   "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
               }
  ### Excluir paciente [DELETE /api/v1/patients/{id}]
  + Request (application/json)
    + Header

             {"SECRET_TOKEN":"TOKEN"}
  + Response 200 (application/json)
    + Body

               {
                   "success": "patient deleted"
               }
  + Response 400 (application/json)
    + Body

               {
                   "code": 400,
                   "status": "error",
                   "message": "invalid id provided",
                   "time_stamp": "2022-12-07 14:55:13"
               }
  + Response 404 (application/json)
    + Body

               {
                   "code": 404,
                   "status": "error",
                   "message": "unable to delete, patient not found",
                   "time_stamp": "2022-12-07 14:55:13"
               }

### Administração de marcação de consultas:

Deve ser possível atribuir a um **paciente** uma consulta com um **dentista** em uma determinada **data e hora**, e também adicionar uma **descrição** à consulta. Solicita-se o desenvolvimento de um _CRUD_ para a entidade **Consulta**.
   
 #### Métodos

 Requisições para a API devem seguir os padrões:

 | Método   | Descrição                                                                                                                |
 |----------|--------------------------------------------------------------------------------------------------------------------------|
 | `POST`   | Adicionar consulta                                                                                                       |
 | `GET`    | Trazer consulta pelo ID                                                                                                  |
 | `PUT`    | Atualizar consulta                                                                                                       |
 | `PATCH`  | Atualizar uma consulta por algum dos seus campos                                                                         |
 | `DELETE` | Excluir consulta.                                                                                                        |
 | `POST`   | Adicionar consulta pelo RG do paciente e matrícula do dentista.                                                          |
 | `GET`    | Trazer consulta pelo RG do paciente. Deve conter o detalhamento da consulsta (Data-Hora, descrição, Paciente e Dentista) |
### Request / Responses
### Adicionar consulta [POST /api/v1/appointments]
+ Request (application/json)
    + Header

            {"SECRET_TOKEN": "TOKEN"}
    + Body

            {
              "description": "Appointment description",
              "date_and_time": "07/01/2023 14:00",
              "dentist_license": "123456",
              "patient_identity": "00000001"
            }
+ Response 201 (application/json)
    + Body

           {
               "data": {
                   "id": 1,
                   "description": "Appointment description",
                   "date_and_time": "07/01/2023 14:00",
                   "dentist_license": "05000001",
                   "patient_identity": "04333330",
                   "dentist": {
                       "id": 1,
                       "surname": "Silva",
                       "name": "João",
                       "license_number": "05000001"
                   },
                   "patient": {
                      "id": 1,
                      "surname": "Queiroz",
                      "name": "Julia",
                      "identity_number": "04333330",
                      "created_at": "07/12/2022 11:00"
                   }
               }
           }
+ Response 400 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "invalid data provided",
               "time_stamp": "2022-12-08 18:42:45.645419 -0300 -03 m=+3.702563901"
           }
+ Response 401 (application/json)
    + Body

           {
               "code": 401,
               "status": "error",
               "message": "Token not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
### Listar consultas [GET /api/v1/appointments]
+ Request (application/json)
    + Header

            {"SECRET_TOKEN": "TOKEN"}
+ Response 200 (application/json)
    + Body

           {
              "data": [
                  {
                      "id": 1,
                      "description": "Appointment description",
                      "date_and_time": "07/01/2023 14:00",
                      "dentist_license": "05000001",
                      "patient_identity": "04333330",
                      "dentist": {
                          "id": 1,
                          "surname": "Silva",
                          "name": "João",
                          "license_number": "05000001"
                      },
                      "patient": {
                          "id": 1,
                          "surname": "Queiroz",
                          "name": "Julia",
                          "identity_number": "04333330",
                          "created_at": "07/12/2022 11:00"
                      }
                  },
                  {
                      "id": 1,
                      "description": "Appointment description",
                      "date_and_time": "07/02/2023 14:00",
                      "dentist_license": "05000001",
                      "patient_identity": "04333330",
                      "dentist": {
                          "id": 1,
                          "surname": "Silva",
                          "name": "João",
                          "license_number": "05000001"
                      },
                      "patient": {
                          "id": 1,
                          "surname": "Queiroz",
                          "name": "Julia",
                          "identity_number": "04333330",
                          "created_at": "07/12/2022 11:00"
                      }
                  }
              ]
           }
+ Response 400 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "invalid request",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 401 (application/json)
    + Body

           {
               "code": 401,
               "status": "error",
               "message": "Token not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
### Listar consultas por dentista [GET /api/v1/appointments/dentist/{license_number}]
+ Request (application/json)
    + Header

            {"SECRET_TOKEN": "TOKEN"}
+ Response 200 (application/json)
    + Body

           {
              "data": [
                  {
                      "id": 1,
                      "description": "Appointment description",
                      "date_and_time": "07/01/2023 14:00",
                      "dentist_license": "05000001",
                      "patient_identity": "04333330",
                      "dentist": {
                          "id": 1,
                          "surname": "Silva",
                          "name": "João",
                          "license_number": "05000001"
                      },
                      "patient": {
                          "id": 1,
                          "surname": "Queiroz",
                          "name": "Julia",
                          "identity_number": "04333330",
                          "created_at": "07/12/2022 11:00"
                      }
                  },
                  {
                      "id": 1,
                      "description": "Appointment description",
                      "date_and_time": "07/02/2023 14:00",
                      "dentist_license": "05000001",
                      "patient_identity": "04333330",
                      "dentist": {
                          "id": 1,
                          "surname": "Silva",
                          "name": "João",
                          "license_number": "05000001"
                      },
                      "patient": {
                          "id": 1,
                          "surname": "Queiroz",
                          "name": "Julia",
                          "identity_number": "04333330",
                          "created_at": "07/12/2022 11:00"
                      }
                  }
              ]
           }
+ Response 400 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "invalid request",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 401 (application/json)
    + Body

           {
               "code": 401,
               "status": "error",
               "message": "Token not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 404 (application/json)
    + Body

           {
               "code": 404,
               "status": "error",
               "message": "appointment not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
### Listar consultas por paciente [GET /api/v1/appointments/patient/{identity_number}]
+ Request (application/json)
    + Header

            {"SECRET_TOKEN": "TOKEN"}
+ Response 200 (application/json)
    + Body

           {
              "data": [
                  {
                      "id": 1,
                      "description": "Appointment description",
                      "date_and_time": "07/01/2023 14:00",
                      "dentist_license": "05000001",
                      "patient_identity": "04333330",
                      "dentist": {
                          "id": 1,
                          "surname": "Silva",
                          "name": "João",
                          "license_number": "05000001"
                      },
                      "patient": {
                          "id": 1,
                          "surname": "Queiroz",
                          "name": "Julia",
                          "identity_number": "04333330",
                          "created_at": "07/12/2022 11:00"
                      }
                  },
                  {
                      "id": 1,
                      "description": "Appointment description",
                      "date_and_time": "07/02/2023 14:00",
                      "dentist_license": "05000001",
                      "patient_identity": "04333330",
                      "dentist": {
                          "id": 1,
                          "surname": "Silva",
                          "name": "João",
                          "license_number": "05000001"
                      },
                      "patient": {
                          "id": 1,
                          "surname": "Queiroz",
                          "name": "Julia",
                          "identity_number": "04333330",
                          "created_at": "07/12/2022 11:00"
                      }
                  }
              ]
           }
+ Response 400 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "invalid request",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 401 (application/json)
    + Body

           {
               "code": 401,
               "status": "error",
               "message": "Token not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 404 (application/json)
    + Body

           {
               "code": 404,
               "status": "error",
               "message": "appointment not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
### Buscar consulta pelo ID [GET /api/v1/appointments/{id}]
+ Request (application/json)
    + Header

            {"SECRET_TOKEN": "TOKEN"}
+ Response 200 (application/json)
    + Body

           {
               "data": {
                   "id": 1,
                   "description": "Appointment description",
                   "date_and_time": "07/01/2023 14:00",
                   "dentist_license": "05000001",
                   "patient_identity": "04333330",
                   "dentist": {
                       "id": 1,
                       "surname": "Silva",
                       "name": "João",
                       "license_number": "05000001"
                   },
                   "patient": {
                      "id": 1,
                      "surname": "Queiroz",
                      "name": "Julia",
                      "identity_number": "04333330",
                      "created_at": "07/12/2022 11:00"
                   }
               }
           }
+ Response 400 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "invalid id provided",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 401 (application/json)
    + Body

           {
               "code": 401,
               "status": "error",
               "message": "Token not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 404 (application/json)
    + Body

           {
               "code": 404,
               "status": "error",
               "message": "appointment not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
### Atualizar consulta [PUT /api/v1/appointments/{id}]
+ Request (application/json)
    + Header

            {"SECRET_TOKEN": "TOKEN"}  
    + Body

               {
                  "id": 1,
                   "description": "A better appointment description",
                   "date_and_time": "08/01/2023 13:00",
                   "dentist_license": "05000001",
                   "patient_identity": "04333330"
               }
+ Response 200 (application/json)
    + Body

           {
               "data": {
                   "id": 1,
                   "description": "A better appointment description",
                   "date_and_time": "08/01/2023 13:00",
                   "dentist_license": "05000001",
                   "patient_identity": "04333330",
                   "dentist": {
                       "id": 1,
                       "surname": "Silva",
                       "name": "João",
                       "license_number": "05000001"
                   },
                   "patient": {
                      "id": 1,
                      "surname": "Queiroz",
                      "name": "Julia",
                      "identity_number": "04333330",
                      "created_at": "07/12/2022 11:00"
                   }
               }
           }
+ Response 400 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "invalid id provided",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 401 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "Token not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 404 (application/json)
    + Body

           {
               "code": 404,
               "status": "error",
               "message": "appointment not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
### Atualizar consulta [PATCH /api/v1/appointment/{id}]
+ Request (application/json)
    + Header

            {"SECRET_TOKEN": "TOKEN"}
    + Body

            {
               "description": "A new better appointment description",
               "date_and_time": "03/01/2023 09:00",
            }
+ Response 200 (application/json)
    + Body

           {
               "data": {
                   "id": 1,
                   "description": "A new better appointment description",
                   "date_and_time": "03/01/2023 09:00",
                   "dentist_license": "05000001",
                   "patient_identity": "04333330",
                   "dentist": {
                       "id": 1,
                       "surname": "Silva",
                       "name": "João",
                       "license_number": "05000001"
                   },
                   "patient": {
                      "id": 1,
                      "surname": "Queiroz",
                      "name": "Julia",
                      "identity_number": "04333330",
                      "created_at": "07/12/2022 11:00"
                   }
               }
           }
+ Response 400 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "invalid id provided",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 400 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "invalid data provided",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 401 (application/json)
    + Body

           {
               "code": 401,
               "status": "error",
               "message": "Token not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 404 (application/json)
    + Body

           {
               "code": 404,
               "status": "error",
               "message": "appointment not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
### Excluir consulta [DELETE /api/v1/appointment/{id}]
+ Request (application/json)
    + Header

          {"SECRET_TOKEN": "TOKEN"}
+ Response 200 (application/json)
    + Body

           {
               "success": "appointment removed"
           }
+ Response 400 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "invalid id provided",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }
+ Response 401 (application/json)
    + Body

          {
            "code": 400,
            "status": "error",
            "message": "Token not found",
            "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
          }
+ Response 404 (application/json)
    + Body

           {
               "code": 400,
               "status": "error",
               "message": "unable to remove, appointment not found",
               "time_stamp": "2022-12-07 14:55:13.645419 -0300 -03 m=+3.702563901"
           }

## Requerimentos técnicos
A aplicação deve ser desenvolvida em design orientado a pacotes:
* **Camada/domínio** de entidades de negócio
* **Camada/domínio** de acesso a dados (Repository)
* **Camada de acesso a dados (banco de dados)**: é o banco de dados do nosso sistema. Você poderá usar qualquer um dos bancos de dados relacionais modelado através de um modelo
entidade-relação, como H2 ou MySQL, ou não relacional, como o Mongo DB.
* **Camada/domínio** service
* **Camada/domínio** handler