# Gestor de Tareas en Go

¡Bienvenido al proyecto del Gestor de Tareas en Go! Este repositorio acompaña al post "NIVEL 3: CREEMOS UN GESTOR DE TAREAS EN GOLANG" de nuestro blog gognition.pro donde aprendimos a crear un gestor de tareas utilizando Go.

## 📋 Descripción

Este proyecto implementa un gestor de tareas en línea de comandos usando Go. Permite añadir, listar y marcar tareas como completadas, además de guardar y cargar tareas desde un archivo JSON. Introduce conceptos fundamentales de programación en Go como structs, métodos, slices y manejo de archivos.

## 🚀 Comenzando

Para utilizar este gestor de tareas, asegúrate de tener Go instalado en tu sistema.

### Prerrequisitos

- Go (versión 1.15 o superior recomendada)

### Instalación

1. Haz un fork de este repositorio haciendo clic en el botón "Fork" en la parte superior derecha de esta página.

2. Clona tu fork a tu máquina local:

```bash
git clone https://github.com/TU_USUARIO/gognition-nivel3-gestor-tareas.git
```

## 💻️ Uso
Para ejecutar el gestor de tareas
```bash
go run main.go
```

El programa te presentará un menú con las siguientes opciones:

1. Añadir tarea
2. Listar tareas
3. Marcar tarea como completada
4. Salir
```bash
1. Añadir tarea
2. Listar tareas
3. Marcar tarea como completada
4. Salir
Elige una opción: 1
Título de la tarea: Hacer ejercicio
Descripción: Correr 5 km
```

## 🏗️ Compilación (Build)
Para compilar el proyecto y crear un ejecutable, puedes usar los siguientes comandos dependiendo de tu sistema operativo objetivo:
```bash
go build -o gestor-tareas
```

Para Windows (desde Linux o macOS):
```bash
GOOS=windows GOARCH=amd64 go build -o gestor-tareas.exe
```

Para macOS (desde Windows o Linux):
```bash
GOOS=darwin GOARCH=amd64 go build -o gestor-tareas-mac
```

Para Linux (desde Windows o macOS):
```bash
GOOS=linux GOARCH=amd64 go build -o gestor-tareas-linux
```

### Después de la compilación, puedes ejecutar el programa usando:
```
./gestor-tareas  # En Linux o macOS
gestor-tareas.exe  # En Windows
```

# Visita gognition.pro https://www.gognition.pro/