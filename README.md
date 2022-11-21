Modelado y Programación
=======================

Proyecto 2: Reproductor de música
-----------------


### DISEÑO:

Realizaremos un diseño orientado a objetos y utilizaremos el patrón de diseño Modelo-Vista-Controlador (MVC).

* Modelo
Para la parte del modelo tendremos los objetos Miner, DataBase, Rola, Search Parser y Builder.
Miner será el encargado de obtener las etiquetas IDV2v3 de los archivos de música en el directorio que ingrese el usuario. Comenzará su recorrido buscando archivos con extensión mp3 en la raíz del directorio de entrada, al encontrar alguno procederá a obtener y guardar las etiquetas junto con la ruta del archivo en un objeto Rola. Si Miner se encuentra con un subdirectorio en su recorrido entonces trabajara recursivamente sobre él. Todas las Rolas serán almacenadas en una colección.
Una vez obtenidas todas las etiquetas Builder creara y poblara la base de datos utilizando la colección de Rolas.
DataBase representará a la base datos, podrá realizar consultas u operaciones de inserción y eliminación.
SearchParser se encargará de convertir las búsquedas del usuario en sentencias SQL que posteriormente serán utilizadas para realizar consultas sobre la base de datos.

* Vista
En la vista tendremos un objeto Window Handler que se ocupara con todo lo relacionado a la creación de las distintas partes de la interfaz gráfica.
La interfaz tendrá botones para saltar a la canción anterior o a la siguiente, botón para silenciar, pausar, reproducir así como una barra que indicara el avance de la canción. Las búsquedas solicitadas por el usuario aparecerán en un display debajo de la barra de búsquedas.
Al iniciar el programa el usuario tendrá que ingresar la ruta del directorio donde tiene guardadas las canciones.
Si en algún momento ocurre un error el usuario será notificado con un mensaje y de ser necesario se terminara la ejecución del programa.

* Controlador
En el controlador tendremos un objeto MainApp el cual tendrá como propiedades un WindowHandler, un DataBase, así como la ruta de directorio ingresada por el usuario. MainApp a través de su WindowHandler construirá la interfaz de usuario y asignará funcionalidades a cada uno de los botones de la misma.

### Compilación
    
Para compilar el programa ejecutar el siguiente comando:
```
$ go build cmd/rmp.go
```
     
### Ejecución

Para ejecutar el programa ejecutar el siguiente comando:
```
$ ./rmp
```
### Uso
Es necesario descargar ALSA. En Ubuntu o Debian, correr el comando:
$ apt install libasound2-dev

En distribuciones de linux RedHat-based:
$ dnf install alsa-lib-devel

La ruta de directorio debe empezar con el caracter "/"

### Pruebas Unitarias
Para correr las pruebas unitarias se deben ejecutar los siguientes comandos:
```
$ go test -v test/miner/*.go
$ go test -v test/database/*.go (es necesario borrar el archivo .db en test/database para volver a correr las pruebas unitarias de la base de datos).
 ```

