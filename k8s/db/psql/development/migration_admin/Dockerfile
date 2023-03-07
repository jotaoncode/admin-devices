FROM migrate/migrate:4

COPY ./sql ./sql
COPY run_migration.sh .
RUN chmod 655 run_migration.sh
USER postgres
