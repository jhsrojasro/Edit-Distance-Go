package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "math"
)

//Matriz de la dp variable global
var distancias = make([][]int, 10000)
var rows = make([]int, 10000*10000)


//variables globales
var s1 ,s2 string
var i, j int



func main() {
    inicializarMatriz()
    fmt.Println("Escriba las palabras reservadas del primer lenguaje en una sola linea separadas por un espacio")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    lenguaje1 := strings.Split(scanner.Text()," ")
    /*
    for i:=0; i<len(lenguaje1); i++{
        fmt.Println(lenguaje1[i])    
    }
    */
    fmt.Println("Escriba las palabras reservadas del segundo lenguaje en una sola linea separadas por un espacio")
    scanner.Scan()
    lenguaje2 := strings.Split(scanner.Text()," ")
    /*
    for i:=0; i<len(lenguaje2); i++{
        fmt.Println(lenguaje2[i])    
    }
    */
    N := len(lenguaje1)
    M := len(lenguaje2)
    distancias_lenguajes := make([][]int, N)
    filas := make([]int, N*M)
    for I := 0; I < N; I++ {
        distancias_lenguajes[I] = filas[I*M : (I+1)*M]
    }
    
    for I:=0; I<N; I++ {
        for J:=0; J<M; J++ {
            distancias_lenguajes[I][J] = levenshtein(lenguaje1[I], lenguaje2[J])
        }
    }
    promedio := calcularPromedio(distancias_lenguajes, N, M)
    varianza := calcularVarianza(distancias_lenguajes, N , M, promedio)
    
    fmt.Println("Promedio: ",promedio)
    fmt.Println("Varianza: ",varianza)
    fmt.Println("Distancia custom: ", custom_distance(lenguaje1, lenguaje2))
    //imprimirMatriz(distancias_lenguajes, N, M)
    
}

func levenshtein(string1 string, string2 string) int{
    s1 = string1
    s2 = string2
    var n, m int
    n = len(s1) + 1
    m = len(s2) + 1
    
    distancias[0][0] = 0
    for i=1; i<n; i++{
        distancias[i][0] = i 
    }
    for j=1; j<m; j++{
        distancias[0][j] = j 
    } 
    
    for i = 1; i<n; i++ {
        for j = 1; j<m; j++ {
            distancias[i][j] = -1
        }
    }
    
    
    return distance(n-1,m-1)
    
}

func custom_distance(lenguaje1 []string, lenguaje2 []string) float64{
    //Conteo de palabras de diferentes longitudes en el lenguaje 1
    count1 := make(map[int] int)
    existe1 := make(map[int] bool)
    for i=0; i<len(lenguaje1); i++ {
        if existe1[len(lenguaje1[i])] {
            count1[len(lenguaje1[i])] += 1
        }else {
            count1[len(lenguaje1[i])] = 1
            existe1[len(lenguaje1[i])] = true
        }
    }
    
    //Conteo de palabras de diferentes longitudes en el lenguaje 2
    count2 := make(map[int] int)
    existe2 := make(map[int] bool)
    for i=0; i<len(lenguaje2); i++ {
        if existe2[len(lenguaje2[i])] {
            count2[len(lenguaje2[i])] += 1
        }else {
            count2[len(lenguaje2[i])] = 1
            existe2[len(lenguaje2[i])] = true
        }
    }
    var c_distance float64 = 0.0
    var tam int = 0
    if len(count1) > len(count2) {
        for key, value := range count1 {
            if existe2[key] {
                c_distance += math.Abs(float64(value - count2[key]))
                tam += 1
            }
        }  
    }else {
        for key, value := range count2 {
            if existe1[key] {
                c_distance += math.Abs(float64(value - count1[key]))
                tam += 1
            }
        }
    }
    return c_distance / float64(tam)
}

func distance(x int, y int) int{
    if distancias[x][y] != -1{
        return distancias[x][y]
    }else{
        distancias[x][y] = Min3(distance(x,y-1) + 1, distance(x-1,y) + 1, distance(x-1, y-1)+cost(x,y))
        return distancias[x][y]
    }
}

func cost(x int, y int) int{
    if s1[x-1] == s2[y-1]{
        return 0
    }else{
        return 1
    }
}

func Min2(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func Min3(x, y , z int) int {
    return Min2(x, Min2(y, z))
}

func inicializarMatriz(){
    for i = 0; i < 10000; i++ {
        distancias[i] = rows[i*10000 : (i+1)*10000]
    }
}

func calcularPromedio(matriz [][]int, n int, m int) float64{
    promedio := 0.0
    for i=0; i<n; i++ {
        for j=0; j<m; j++ {
            promedio += float64(matriz[i][j])
        }
    }
    return promedio / float64(n*m)
}

func calcularVarianza(matriz [][]int, n int, m int, promedio float64) float64{
    varianza := 0.0
    for i=0;i<n; i++ {
        for j=0; j<m; j++ {
            varianza += (float64(matriz[i][j]) - promedio) * (float64(matriz[i][j]) - promedio) 
        }
    }
    return varianza / float64(n*m)
}

func imprimirMatriz(matriz [][]int, n int, m int){
    var i, j int
    for i = 0; i<n; i++ {
        for j = 0; j<m; j++ {
            fmt.Print(matriz[i][j]," ")
        }
        fmt.Println()
    }
}