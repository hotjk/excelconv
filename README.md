# excelcov

Convert one excel file to another excel file based on template

## How to use

~~~
excelcov template.json Book1.xlsx Book2.xlsx
~~~

## Template

~~~
{
  "fromIndex": 1,   // source file sheet index, start from 1
  "toIndex": 1,     // target file sheet index, start from 1
  "links": [        // cell to cell
    {
      "to": {
        "row": 3,   // position cell with index, start from 1
        "column": 7
      },
      "from": [
        {
          "name": "G3", // position cell with name
          "func" : "ToUpper"  // function apply to source value after read
        }
      ]
    },
    {
      "to": {
        "name":"G4",
        "func" : "ToUpper"  // function apply to target value before write
      },
      "from": [
        {
          "name": "G4"
        }
      ]
    },
    {
      "to": {
        "name":"G2"
      },
      "from": [
        {
          "name": "G2",
          "func" : "TimeFormat",
          "params" : ["Jan 02, 2006", "2006-01-02"] // function parameters
        }
      ]
    }
  ],
  "loops": [  // repeat copy row
    {
      "stop": {
        "name": "D6"  // break repeat when the value of this cell is empty
      },
      "links": [
        {
          "to": {
            "name": "A1"
          },
          "from": [
            {
              "name": "D6",
              "func": "TimeFormat",
              "params" : ["Jan 02, 2006", "2006-01-02"]
            }
          ]
        }
      ]
    },
    {
      "stop": {
        "name": "A3"
      },
      "links": [
        {
          "to": {
            "name": "E2"
          },
          "from": [
            {
              "name": "A3"
            }
          ]
        },
        {
          "to": {
            "name": "F2"
          },
          "from": [
            {
              "name": "B3"
            },
            {
              "name": "C3"  //combine multiple cell values into one target cell
            }
          ]
        }
      ]
    },
    {
      "stop": {
        "name": "A3"
      },
      "links": [
        {
          "to": {
            "name": "H2"
          },
          "from": [
            {
              "name": "A3"
            },
            {
              "value": "-"  // direct value
            },
            {
              "name": "A3", // will be ignore
              "func":"Sequence",
              "params":["test", "%06d"]
            }
          ]
        },
        {
          "to": {
            "name": "I2"
          },
          "from": [
            {
              "name": "$B3" // start with $, fixed cell, will not increase the row number in repeat
            },
            {
              "value":"----"
            },
            {
              "name": "C3"
            }
          ]
        }
      ]
    }
  ]
}
~~~

## Extention

Try to add your custom function in functions.go
