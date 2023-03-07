# Migrations

### Create migration
To create a new migration just run the sh file for creation like this:

`sh create_migration.sh [NAME_MIGRATION]`

this will create the given migration files with the time stamp, up and down.

### Run the migration

Execute the given migration with the corresponding command.

Why we change the command each time inside of this run_migration.yml job?

I want to trace with the given commits each time we make a migration in github.

This way we have a history of the given migrations in the system.