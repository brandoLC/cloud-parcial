from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware   # ← Importar CORS
import httpx
import os

app = FastAPI()

# 1) Añadir el middleware de CORS
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],            # permite cualquier origen
    allow_credentials=True,
    allow_methods=["*"],            # permite todos los métodos (GET, POST, etc)
    allow_headers=["*"],            # permite todas las cabeceras
)

USUARIOS_SERVICE_URL = os.getenv("USUARIOS_SERVICE_URL", "http://mi-nestjs-container:3000")
CURSOS_SERVICE_URL   = os.getenv("CURSOS_SERVICE_URL",   "http://mi-fiber-container:8070")
NOTAS_SERVICE_URL    = os.getenv("NOTAS_SERVICE_URL",    "http://mi-fastapi-container:8080")

@app.get("/estudiante/{id_estudiante}/cursos")
async def obtener_estudiante_cursos(id_estudiante: int):
    cursos_url = f"{CURSOS_SERVICE_URL}/estudiante-cursos/estudiante/{id_estudiante}"
    async with httpx.AsyncClient() as client:
        cursos_response = await client.get(cursos_url)
        if cursos_response.status_code != 200:
            raise HTTPException(status_code=500, detail="Error al obtener cursos del estudiante")

        cursos_inscritos = cursos_response.json()
        detalles_cursos = []

        for curso in cursos_inscritos:
            codigo_curso = curso["CursoCodigo"]
            curso_detalle_url = f"{CURSOS_SERVICE_URL}/cursos/{codigo_curso}"
            curso_response = await client.get(curso_detalle_url)
            if curso_response.status_code == 200:
                detalles_cursos.append(curso_response.json())
            else:
                detalles_cursos.append({
                    "codigo": codigo_curso,
                    "error": "Información del curso no disponible"
                })

    return detalles_cursos

@app.get("/estudiante/{id_estudiante}/curso/{codigo_curso}/notas")
async def obtener_notas_curso_estudiante(id_estudiante: int, codigo_curso: int):
    notas_url = f"{NOTAS_SERVICE_URL}/notas"
    params = {"id_estudiante": id_estudiante, "codigo_curso": codigo_curso}
    async with httpx.AsyncClient() as client:
        notas_response = await client.get(notas_url, params=params)
        if notas_response.status_code != 200:
            raise HTTPException(
                status_code=404,
                detail="No se encontraron notas para este estudiante en el curso especificado"
            )
        notas = notas_response.json()
    return notas
