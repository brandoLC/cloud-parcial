# Usa la imagen oficial de Python
FROM python:3.11-slim

# Crear el directorio de trabajo
WORKDIR /app

# Copiar y instalar dependencias
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copiar el script de seed, el .env y el helper
COPY seed.py .env wait-for-it.sh ./
RUN chmod +x wait-for-it.sh

# ENTRYPOINT: primero espera a cada servicio, luego ejecuta seed.py
ENTRYPOINT ["./wait-for-it.sh", "postgres-db:5432", "--", "./wait-for-it.sh", "mysql-db:3306", "--", "./wait-for-it.sh", "mongo-db:27017", "--", "python", "seed.py"]
