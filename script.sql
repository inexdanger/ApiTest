CREATE TABLE preguntas (
  IdPregunta INT NOT NULL AUTO_INCREMENT,
  IdExamen VARCHAR(255) NOT NULL,
  Pregunta VARCHAR(255),
  Respuesta1 VARCHAR(255),
  Respuesta2 VARCHAR(255),
  Respuesta3 VARCHAR(255),
  Respuesta4 VARCHAR(255),
  Correcta VARCHAR(255),
  PRIMARY KEY (IdPregunta)
);

CREATE TABLE examenes (
  IdExamen VARCHAR(255) NOT NULL PRIMARY KEY,
  CantidadPreguntas INT
);