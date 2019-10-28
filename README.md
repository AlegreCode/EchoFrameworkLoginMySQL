# EchoFrameworkLoginMySQL
### Plantilla de inicio de sesión y registro de usuario con Echo framework y MySQL.

##### Packages utilizados:
- **Sprig**: extiende las funcionalidades de template de golang. [Aquí](https://github.com/Masterminds/sprig)
- **gookit/validate**: agrega funciones de validación de entradas. [Aquí](https://github.com/gookit/validate)
- **godotenv**: permite cargar variables de entorno desde un archivo .env. [Aquí](https://github.com/joho/godotenv)
- **bcrypt**: utilizado para encriptar las contraseñas de los usuarios. [Aquí](https://godoc.org/golang.org/x/crypto/bcrypt)
- **go-sql-driver/mysql**: agrega el driver para conexiones mysql. [Aquí](http://go-database-sql.org/index.html)

##### Config base de datos
El archivo .env-example renombrar a .env y completar las variables con tus credenciales de conexión.
