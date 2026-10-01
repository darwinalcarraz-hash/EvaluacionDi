package main

import "fmt"

// Variables globales (slices)
var nombresProductos []string
var subtotales []float64

// Función para registrar una venta
func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)
	nombresProductos = append(nombresProductos, nombre)
	subtotales = append(subtotales, subtotal)
	fmt.Println("Venta registrada con éxito.\n")
}

// Función para mostrar estadísticas
func MostrarEstadisticas() {
	if len(nombresProductos) == 0 {
		fmt.Println("No existen ventas registradas.\n")
		return
	}

	totalRecaudado := 0.0
	fmt.Println("\n--- ESTADÍSTICAS ---")
	for i := 0; i < len(nombresProductos); i++ {
		fmt.Printf("Producto: %s | Subtotal: $%.2f\n", nombresProductos[i], subtotales[i])
		totalRecaudado += subtotales[i]
	}
	fmt.Printf("Total recaudado: $%.2f\n", totalRecaudado)
}

func main() {
	var opcion int

	for {
		fmt.Println("MENU PRINCIPAL")
		fmt.Println("1. Registrar una nueva venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			fmt.Println("\nLista de productos:")
			fmt.Println("1. Arroz ($1.25)")
			fmt.Println("2. Leche ($0.95)")
			fmt.Println("3. Pan ($0.50)")
			fmt.Print("Seleccione el número del producto: ")

			var prod int
			fmt.Scan(&prod)

			fmt.Print("Ingrese la cantidad vendida: ")
			var cantidad int
			fmt.Scan(&cantidad)

			// Seleccionar precio y llamar a la función
			switch prod {
			case 1:
				RegistrarVenta("Arroz", 1.25, cantidad)
			case 2:
				RegistrarVenta("Leche", 0.95, cantidad)
			case 3:
				RegistrarVenta("Pan", 0.50, cantidad)
			default:
				fmt.Println("Número de producto inválido.\n")
			}

		case 2:
			MostrarEstadisticas()

		case 3:
			fmt.Println("¡Hasta luego!")
			return

		default:
			fmt.Println("Opción no válida.\n")
		}
	}
}
