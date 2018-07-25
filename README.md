# ine-callejero-dataset

Este proyecto contiene un listado completo del callejero proporcionado por el Instituto Nacional de Estadística (INE).

El archivo original puede encontrarse en la siguiente URL:
- [Callejero del Censo Electoral (INE)](http://www.ine.es/ss/Satellite?L=es_ES&c=Page&cid=1254735624326&p=1254735624326&pagename=ProductosYServicios%2FPYSLayout)

Desde esta URL puedes descargarte dos archivos por cada año a nivel nacional. La URL directa a los archivos es:
- `http://www.ine.es/prodyser/callejero/caj_esp/caj_esp_[MM][YYYY].zip`
  - Sustituye `MM` por el mes que quieras (01 para enero ó 07 para julio)
  - Sustituye `YYYY` por el año que quieras

Este proyecto contiene dos carpetas diferentes:
- Dentro de la carpeta [/archive](/archive) encontrarás el último fichero *.zip* correspondiente a la fecha más cercana del último commit.
- Dentro de la carpeta [/data](/data) encontrarás los ficheros procesados del archivo correspondiente al punto anterior en formato *.json* para mayor comodidad.
  - *NOTA*: GitHub no permite subir archivos de más de 100 Mb por lo que podrás encontrarlos en la sección de [Releases](https://github.com/jontorrado/ine-callejero-dataset/releases)


## Archivos descargados

Al descomprimir el archivo *.zip* descargado, nos encontraremos con 4 archivos (al menos) diferentes:
- PSEU-NAL: pseudovías nacionales
- TRAMOS-NAL: tramos nacionales
- UP-NAL: unidades poblacionales nacionales
- VIAS-NAL: vías nacionales

El formato de los archivos lo tienes explicado en [este documento Word](http://www.ine.es/ss/Satellite?blobcol=urldata&blobheader=application%2Fmsword&blobheadername1=Content-Disposition&blobheadervalue1=attachment%3B+filename%3Ddisenocallejero.doc&blobkey=urldata&blobtable=MungoBlobs&blobwhere=817%2F26%2Fdisenocallejero%2C1.doc&ssbinary=true) del INE.

Para ver los campos de los diferentes modelos, mira el código de la carpeta [/models](/models).

## Organizaciones administrativas

Los datos de comunidades, provincias, municipios e islas los puedes encontrar en el siguiente repositorio:
- [https://github.com/codeforspain/ds-organizacion-administrativa](https://github.com/codeforspain/ds-organizacion-administrativa)

# Ejecutable de GO
Se proporciona un ejecutable dentro de la carpeta [/build](/build) para poder crear los archivos *.json* actualizados. Basta con escribir en la consola `./build/ine-callejero-dataset` y el programa te solicitará el nombre (con el path) al archivo. Si descomprimes el archivo *.zip* de la carpeta [/archive](/archive) dentro de build, solo tendrías que escribir `PSEU-NAL.F180312` para que procese dicho archivo correspondiente a las pseudovias nacionales.

### Agradecimientos
Agradecimiento a Iñigo Flores por su repositorio [https://github.com/inigoflores/ds-codigos-postales-ine-es](https://github.com/inigoflores/ds-codigos-postales-ine-es)
