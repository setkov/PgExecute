# PgExucute

Tool for execute sql scripts in Postgre Sql database:
1. connect to the database
2. sequentially (in alphabetical order) execute sql scripts from the input folder
3. the executed script is transferred to the output folder
4. if an error occurs, file processing stops

Set the settings in the configuration file, see the pgexecute-sample.yaml file for an example:
- _connection:_ - connection string
- _files:_
  - _in:_ - input forder
  - _out:_ - output folder

Run:
- ```pgexecute.exe``` - run with default configuration file name pgexecute.yaml
- ```pgexecute.exe other.yaml``` - run with other configuration file name