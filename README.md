# lem-in

This project is a digital version of an ant farm simulation. The goal of the project is to create a program called `lem-in` that can efficiently navigate ants through a colony by finding the shortest path from the start room to the end room. The project involves reading and parsing input files, implementing the Edmonds-Karp algorithm, and displaying the movements of the ants.

## Features

- Reads the ant colony description from a file
- Implements the Edmonds-Karp algorithm to find the shortest path
- Displays the content of the input file
- Displays each move the ants make from room to room
- Handles invalid or poorly formatted input data gracefully
- Provides a clean and well-structured codebase

## Getting Started

To get started with the `lem-in` project, follow these steps:

1. Clone the repository:
   ```
   git clone <repository_url>
   ```

2. Compile the program:
   ```
   go build
   ```

3. Run the program with an input file:
   ```
   ./lem-in <input_file>
   ```

   Replace `<input_file>` with the path to your input file.

4. Observe the program's output, which includes the content of the input file and the movements of the ants.

## Input File Format

The input file should follow a specific format. Here is an example:

```
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1
```

The input file consists of the following elements:

- Room definitions: Each room is defined as `<name> <coord_x> <coord_y>`. The format is flexible, and the room names can be alphanumeric.
- Start and end room definitions: The start room is marked with `##start`, and the end room is marked with `##end`.
- Tunnel definitions: Each tunnel is defined as `<room_name1>-<room_name2>`. The tunnels connect the rooms in the colony.

## Output Format

The program outputs the following information:

1. The number of ants (`number_of_ants`)
2. The room definitions (`the_rooms`)
3. The tunnel definitions (`the_links`)
4. The movements of the ants in each turn (`Lx-y Lz-w Lr-o ...`)

The movements of the ants are represented as `Lx-y`, where `x` is the ant number and `y` is the room name.

```
L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3 L4-2
L3-1 L4-3
L4-1
```

## Error Handling

The program is designed to handle various error scenarios gracefully. If the input data is invalid or poorly formatted, the program will display an error message indicating the specific issue. For example, it may show an error message like:

```
ERROR: invalid data format
```

or

```
ERROR: invalid number of ants
```

Make sure to check the error messages and adjust the input file accordingly.

## Contributing

Contributions are welcome! If you have any ideas, improvements, or bug fixes, please submit a pull request. For major changes, please open an issue first to discuss the proposed changes. Feel free to explore, experiment, and improve the `lem-in` program. Enjoy the challenge of navigating ants through a digital ant farm!

## Acknowledgments

This project was completed as part of a programming challenge at [Alem School](https://alem.school).

---
	|NNNNNNNNNNNNNNWNXKK000KKXNNWNNNNNNNNNNNNNNNNNNNNNN|     |
	|NNNNNNNNNNNKxoc,'........',:lx0NNWk;''''''''''''lX|     |	
	|NNNNNNNN0o;.                  .,lOd.            ;X|     |  
	|NNNNNXk:.                        .,,            ;X|     |  
	|NNNNO;                            ,Ox'          ;X|     |  
	|NNXd.                             ,KWKc         ;X|     |  
	|NXo.                              ,KWNXo.       ;X|     |	
	|Nd.                               ,KWNNXc       ;X|     |	
	|0'                                ,KWNNW0'      ;X|     |  
	|d                                 ,KWNNNNc      ;X|     |  
	|c                                 ,KWNNNNd      ;X|     |  
	|c                                 ,KWNNNNo      ;X|     |  
	|d.                                ,KWNNNXc      ;X|     |  
	|0,                                ,KWNNWO'      ;X|     |  
	|Nd.                               ,KWNNX:       ;X|     |  
	|NNo.                              ,KWNXl        ;X|     |  
	|NNNx.                             ,KW0:         ;X|     |	Made by amayev && HgCl2
	|NNNN0:.                           ,Od.          ;X|     |  https://alem.school ...
	|NNNNNNOc.                        .,'            ;X|     |
	|NNNNNNNN0d:.                  .;o0d.            ;X|     |
	|NNNNNNNNNNNXOdc;,.......',;cokKNNNOc::::::::::::dX|     |
	|NNNNNNNNNNNNNNWNXKK000KKXNNWNNNNNNNNNNNNNNNNNNNNNN|     |

## Collaborators
---
 amayev - https://github.com/ive663
 HgCl2  - https://github.com/HgCl2
## License
---

This project is licensed under the [MIT License](LICENSE).
```
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, 
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF 
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. 
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, 
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR 
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR 
THE USE OR OTHER DEALINGS IN THE SOFTWARE. 

Copyright (c) 2023
```