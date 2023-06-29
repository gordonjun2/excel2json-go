# Convert .xlsx or .csv to JSON

## Credit
- Code originally from [excel2json by FerdinaKusumah](https://github.com/FerdinaKusumah/excel2json)
- Adapted to my own preference instead:
	- Set up a proper project directory where the main executable packages are in the ```cmd``` folder and the 
	  helper functions are in ```pkg``` folder
	- Set file path and sheet name in command argument
	- Print array of map instead of map line-by-line
	- Removed remote URL functionality (I do not use it)

## Usage
### Suppose you have Excel table like this
![Excel Image](https://raw.githubusercontent.com/FerdinaKusumah/excel2json/master/image/excel_image.png)

### and you need to convert it to this
```json
[{
    "Profit": "-213.25",
    "ShippingCost": "35",
    "UnitPrice": "38.94"
}, {
    "Profit": "457.81",
    "ShippingCost": "68.02",
    "UnitPrice": "208.16"
}, {
    "Profit": "46.71",
    "ShippingCost": "2.99",
    "UnitPrice": "8.69"
}]
```

### and you have an Excel `.xlsx`, you should do
```
go build ./cmd/xlsxlocal

./xlsxlocal <path to .xlsx file> <.xlsx sheet name, default is the first sheet>
eg.
./xlsxlocal ./samplefiles/iris_dataset.xlsx
```

### and you have an Excel `.csv`, you should do
```
go build ./cmd/csvlocal

./csvlocal <path to .csv file>
eg.
./csvlocal ./samplefiles/iris_dataset.csv
```
