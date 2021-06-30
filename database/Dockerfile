FROM postgres:13.3

ENV POSTGRES_USER=golinksadm
ENV POSTGRES_PASSWORD=mdasknilog
ENV POSTGRES_DB=golinks

ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY initialize.sql /docker-entrypoint-initdb.d/initialize.sql

EXPOSE 5432