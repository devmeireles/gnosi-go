# Gnosi API

Gnosi is a tool capable of connecting tutors and students in a simple and objective way, with it you can improve various skills according to your desire to learn and of course, to teach. Gnosi is a digital space that welcomes students, teachers, curious and enthusiastic.
Without a doubt this is the most important part of the project, here is the heart, the nervous system and maybe the whole brain of our operation.

## [Installation](#installation)

- [Installation](#installation)
- [Configurations](#configurations)
- [Step-by-step](#step-by-step)
  - [Creating a Feature](#creating-a-feature)

Use the package manager [npm](http://npmjs.com/) to install the necessary libraries

```bash
npm i
```

## [Configurations](#configurations)

Make sure to fill in the environment variables in the `.env` file

```python
PORT=
NODE_ENV=
STORAGE_TYPE=
SWAGGER_URL=
BUCKET_NAME=
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_DEFAULT_REGION=

PROD_DB_USERNAME=
PROD_DB_PASSWORD=
PROD_DB_NAME=
PROD_DB_HOST=
PROD_DB_URL=
PROD_DB_DIALECT=
PROD_DB_PROTOCOL=
```

## [Step-by-step](#step-by-step)

The development of Gnosi is based on [TDD](https://en.wikipedia.org/wiki/Test-driven_development) to guarantee an effective development, for this, some steps must be followed strictly to maintain a development pattern. In addition, the entire structure is divided into controllers, models and services.

### [Creating a feature](#creating-a-feature)

To build a feature you need to follow these steps:

**1.** Create the `.test.js` file in the `tests` folder

**2.** Create the router file in the `src/app/routes` folder and use it in the `src/app.js` file

**3.** Create the validation file in the `src/app/middlewares/validator` folder and use it in your created route file in the `src/app/routes` folder

**4.** Create the controller file in the `src/app/controllers` folder and use it in your created route file in the `src/app/routes` folder

**5.** Create the service file in the `src/app/services` folder and use it in your created controller file in the `src/app/controllers` folder

### Folder Structure

```
📦client-api
 ┣ 📂src
 ┃ ┣ 📂app
 ┃ ┃ ┣ 📂config
 ┃ ┃ ┣ 📂controllers
 ┃ ┃ ┣ 📂helpers
 ┃ ┃ ┣ 📂middlewares
 ┃ ┃ ┃ ┗ 📂validator
 ┃ ┃ ┣ 📂migrations
 ┃ ┃ ┣ 📂models
 ┃ ┃ ┣ 📂routes
 ┃ ┃ ┣ 📂seeders
 ┃ ┃ ┗ 📂services
 ┃ ┣ 📜app.js
 ┃ ┗ 📜server.js
 ┣ 📂tests
 ┃ ┣ 📂util
 ┣ 📜.env
 ┣ 📜.eslintrc.js
 ┣ 📜.gitignore
 ┣ 📜.sequelizerc
 ┣ 📜README.md
 ┗ 📜package.json
```

## Routes

| Name             |  Verb  | Address                       | Handler                              |
| ---------------- | :----: | ----------------------------- | ------------------------------------ |
|                  |        | **Catalogue**                 |
| Enum             |  GET   | api/catalogue                 | CatalogueController.index            |
| Read             |  GET   | api/catalogue/{id}            | CatalogueController.read             |
| Create           |  POST  | api/catalogue                 | CatalogueController.create           |
| Update           |  PUT   | api/catalogue/{id}            | CatalogueController.update           |
| Delete           | DELETE | api/catalogue/{id}            | CatalogueController.delete           |
| Add Languages    |  POST  | api/catalogue/{id}/language   | CatalogueController.addLanguage      |
| Add Objective    |  POST  | api/catalogue/{id}/objective  | CatalogueController.addObjective     |
| Add Categories   |  POST  | api/catalogue/{id}/categories | CatalogueController.addCategories    |
| Remove Language  | DELETE | api/catalogue/language/{id}   | CatalogueController.removeLanguage   |
| Remove Objective | DELETE | api/catalogue/objective/{id}  | CatalogueController.removeObjective  |
| Remove Category  | DELETE | api/catalogue/category/{id}   | CatalogueController.removeCategories |
|                  |        | **Season**                    |
| Read             |  GET   | api/category/{id}             | SeasonController.read                |
| Create           |  POST  | api/category                  | SeasonController.create              |
| Update           |  PUT   | api/category/{id}             | SeasonController.update              |
| Delete           | DELETE | api/category/{id}             | SeasonController.delete              |
|                  |        | **Episode**                   |
| Read             |  GET   | api/episode/{id}              | EpisodeController.read               |
| Create           |  POST  | api/episode                   | EpisodeController.create             |
| Update           |  PUT   | api/episode/{id}              | EpisodeController.update             |
| Delete           | DELETE | api/episode/{id}              | EpisodeController.delete             |
|                  |        | **Category**                  |
| Enum             |  GET   | api/category                  | CategoryController.index             |
| Read             |  GET   | api/category/{id}             | CategoryController.read              |
| Create           |  POST  | api/category                  | CategoryController.create            |
| Update           |  PUT   | api/category/{id}             | CategoryController.update            |
| Delete           | DELETE | api/category/{id}             | CategoryController.delete            |
|                  |        | **Language**                  |
| Enum             |  GET   | api/languange                 | LanguageController.index             |
| Read             |  GET   | api/languange/{id}            | LanguageController.read              |
| Create           |  POST  | api/languange                 | LanguageController.create            |
| Update           |  PUT   | api/languange/{id}            | LanguageController.update            |
| Delete           | DELETE | api/languange/{id}            | LanguageController.delete            |

## License

[MIT](https://choosealicense.com/licenses/mit/)