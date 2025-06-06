//Certainly! Here's the modified version of the 3-D N-body simulation in Go, using goroutines and channels for parallelism:
//This version uses goroutines and a sync.WaitGroup to parallelize the force calculations and position updates. The sync.Mutex is used to safely update the shared energy variable during the energy calculation.

package main

import (
	"fmt"
	"math"
	"sync"
)

const (
	G         = 6.67430e-11            // Gravitational constant
	DT        = 1e-3 * 3600 * 24 * 365 // Time step
	NumBodies = 10000                  // Number of bodies
)

type Body struct {
	x, y, z    float64
	vx, vy, vz float64
	mass       float64
}

func initializeBodies(bodies []Body) {
	bodies[0] = Body{0, 0, 0, 0, 0, 0, 1e30} // Central body mass

	for i := 1; i < len(bodies); i++ {
		angle := 2 * math.Pi * float64(i) / float64(len(bodies)-1)
		bodies[i] = Body{
			x:    math.Cos(angle) * 1e11,
			y:    math.Sin(angle) * 1e11,
			z:    0,
			vx:   -math.Sin(angle) * math.Sqrt(G*bodies[0].mass/1e11),
			vy:   math.Cos(angle) * math.Sqrt(G*bodies[0].mass/1e11),
			vz:   0,
			mass: 1e24 / float64(len(bodies)), // Small body mass
		}
	}
}

// func calculateEnergy(bodies []Body) float64 {
// 	energy := 0.0
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex

// 	for i := range bodies {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			kinetic := 0.5 * bodies[i].mass * (bodies[i].vx*bodies[i].vx + bodies[i].vy*bodies[i].vy + bodies[i].vz*bodies[i].vz)
// 			potential := 0.0
// 			for j := range bodies {
// 				if i != j {
// 					dx := bodies[i].x - bodies[j].x
// 					dy := bodies[i].y - bodies[j].y
// 					dz := bodies[i].z - bodies[j].z
// 					distance := math.Sqrt(dx*dx + dy*dy + dz*dz)
// 					potential -= G * bodies[i].mass * bodies[j].mass / distance
// 				}
// 			}
// 			mu.Lock()
// 			energy += kinetic + 0.5*potential
// 			mu.Unlock()
// 		}(i)
// 	}
// 	wg.Wait()
// 	return energy
// }

func kickStep(bodies []Body) {
	var wg sync.WaitGroup

	for i := range bodies {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ax, ay, az := 0.0, 0.0, 0.0
			for j := range bodies {
				if i != j {
					dx := bodies[j].x - bodies[i].x
					dy := bodies[j].y - bodies[i].y
					dz := bodies[j].z - bodies[i].z
					distance := math.Sqrt(dx*dx + dy*dy + dz*dz)
					force := G * bodies[j].mass / (distance * distance * distance)
					ax += force * dx
					ay += force * dy
					az += force * dz
				}
			}
			bodies[i].vx += ax * DT
			bodies[i].vy += ay * DT
			bodies[i].vz += az * DT
		}(i)
	}
	wg.Wait()

	for i := range bodies {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			bodies[i].x += bodies[i].vx * DT
			bodies[i].y += bodies[i].vy * DT
			bodies[i].z += bodies[i].vz * DT
		}(i)
	}
	wg.Wait()
}

func main() {
	bodies := make([]Body, NumBodies)
	initializeBodies(bodies)

	// initialEnergy := calculateEnergy(bodies)
	// fmt.Printf("Initial energy: %e\n", initialEnergy)

	for step := 0; step < 100; step++ {
		kickStep(bodies)
	}

	// finalEnergy := calculateEnergy(bodies)
	// fmt.Printf("Final energy: %e\n", finalEnergy)
	fmt.Printf("Body[0] %e %e %e", bodies[0].x, bodies[0].y, bodies[0].z)

}
