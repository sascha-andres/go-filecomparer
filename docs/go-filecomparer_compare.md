## go-filecomparer compare

Compare files to file database

### Synopsis


Scans the directory and compares to data in file database.

For each change a line will be printed:

D file - file was deleted [not yet implenented]
A file - file was added
C file - file was changed

For an unchanged file, nothing will be printed

```
go-filecomparer compare
```

### Options

```
  -i, --id-only   Just show the paths, not the flag
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.go-filecomparer.yaml)
  -v, --verbose         Turn on verbose output
```

### SEE ALSO
* [go-filecomparer](go-filecomparer.md)	 - A brief description of your application

###### Auto generated by spf13/cobra on 3-Mar-2017
