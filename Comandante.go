package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	pb "github.com/yojeje/lab6"

	"google.golang.org/grpc"
)


func Consultar(m pb.KaisClient) {
	var base1, sector1 string
	scanner := bufio.NewScanner(os.Stdin)

	// Leer el sector
	fmt.Print("Escriba el sector: ")
	if !scanner.Scan() {
		log.Fatalf("Error leyendo el sector: %v", scanner.Err())
	}
	sector1 = scanner.Text()
	sector1 = strings.TrimSpace(sector1)
	if sector1 == "" {
		fmt.Println("El sector no puede estar vacío.")
		return
	}

	// Leer la base
	fmt.Print("Escriba la base: ")
	if !scanner.Scan() {
		log.Fatalf("Error leyendo la base: %v", scanner.Err())
	}
	base1 = scanner.Text()
	base1 = strings.TrimSpace(base1)
	if base1 == "" {
		fmt.Println("La base no puede estar vacía.")
		return
	}

	// Imprimir valores leídos para verificar
	fmt.Println("Sector:", sector1)
	fmt.Println("Base:", base1)

	// Llamar a la función para obtener enemigos, ajusta según tu código real
	response, err := m.GetEnemigosBroker(context.Background(), &pb.Informacion{Base: base1, Sector: sector1})

	if err != nil {
		log.Fatalf("error: %v",err)

	}
	// Establecer conexión con el servidor gRPC
	conn, err := grpc.Dial(response.Dir, grpc.WithInsecure());
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}


	s := pb.NewKaisClient(conn)
	fmt.Println("Enviando comando al servidor" + response.Dir)
	responseS, err := s.GetEnemigosServidor(context.Background(), &pb.Direccion{Dir: response.Dir})

	if err != nil {
		log.Println("ERROR")
	}
	
	fmt.Println(responseS)
}

func main() {
	// Establecer conexión con el servidor gRPC
	conn, err := grpc.Dial("dist098:50051", grpc.WithInsecure());
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	m := pb.NewKaisClient(conn)
	defer conn.Close()

	Consultar(m)
}