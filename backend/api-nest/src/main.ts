import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ValidationPipe } from '@nestjs/common';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  // Habilita CORS para cualquier origen
  app.enableCors({
    origin: '*',
    methods: ['GET','POST','PUT','PATCH','DELETE','OPTIONS'],
    allowedHeaders: ['Content-Type', 'Authorization'],
  });

  // Valida y filtra el payload según tus DTOs
  app.useGlobalPipes(new ValidationPipe({ whitelist: true }));

  const port = Number(process.env.PORT) || 3000;
  // Escucha en todas las interfaces del contenedor
  await app.listen(port, '0.0.0.0');
  console.log(`🚀 Server listening on http://0.0.0.0:${port}`);
}

bootstrap();
