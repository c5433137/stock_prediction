package logic

import "math"

/*
 * TinyEKF: Extended Kalman Filter for embedded processors
 *
 * Copyright (C) 2015 Simon D. Levy
 *
 * MIT License
 */

/* Cholesky-decomposition matrix-inversion code, adapated from
   http://jean-pierre.moreau.pagesperso-orange.fr/Cplus/choles_cpp.txt */
func choldc1(a, p []float64, n int) int {
	i, j, k := 0, 0, 0
	sum := float64(0.0)
	for i = 0; i < n; i++ {
		for j = i; j < n; j++ {
			sum = a[i*n+j]
			for k = i - 1; k >= 0; k-- {
				sum -= a[i*n+k] * a[j*n+k]
			}
			if i == j {
				if sum <= 0 {
					return 1 /* error */
				}
				p[i] = math.Sqrt(sum)
			} else {
				a[j*n+i] = sum / p[i]
			}
		}
	}
	return 0 /* success */
}

func choldcsl(A, a, p []float64, n int) int {
	i, j, k := 0, 0, 0
	sum := float64(0.0)
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			a[i*n+j] = A[i*n+j]
		}
	}
	if choldc1(a, p, n) != 0 {
		return 1
	}
	for i = 0; i < n; i++ {
		a[i*n+i] = 1 / p[i]
		for j = i + 1; j < n; j++ {
			sum = 0
			for k = i; k < j; k++ {
				sum -= a[j*n+k] * a[k*n+i]
			}
			a[j*n+i] = sum / p[j]
		}
	}

	return 0 /* success */
}

func cholsl(A, a, p []float64, n int) int {
	i, j, k := 0, 0, 0
	if choldcsl(A, a, p, n) != 0 {
		return 1
	}
	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			a[i*n+j] = 0.0
		}
	}
	for i = 0; i < n; i++ {
		a[i*n+i] *= a[i*n+i]
		for k = i + 1; k < n; k++ {
			a[i*n+i] += a[k*n+i] * a[k*n+i]
		}
		for j = i + 1; j < n; j++ {
			for k = j; k < n; k++ {
				a[i*n+j] += a[k*n+i] * a[k*n+j]
			}
		}
	}
	for i = 0; i < n; i++ {
		for j = 0; j < i; j++ {
			a[i*n+j] = a[j*n+i]
		}
	}
	return 0 /* success */
}

func zeros(a []float64, m, n int) int {
	j := 0
	for j = 0; j < m*n; j++ {
		a[j] = 0
	}
	return 0
}
func _malloc(m, n int) []float64 {
	return make([]float64,m*n)
}
//#ifdef DEBUG
//static void dump(double * a, int m, int n, const char * fmt)
//{
//int i,j;
//
//char f[100];
//sprintf(f, "%s ", fmt);
//for(i=0; i<m; ++i) {
//for(j=0; j<n; ++j)
//printf(f, a[i*n+j]);
//printf("\n");
//}
//}
//#endif

/* C <- A * B */
func mulmat(a, b, c []float64, arows, acols, bcols int) {
	i, j, l := 0, 0, 0
	for i = 0; i < arows; i++ {
		for j = 0; j < bcols; j++ {
			c[i*bcols+j] = 0
			for l = 0; l < acols; l++ {
				c[i*bcols+j] += a[i*acols+l] * b[l*bcols+j]
			}
		}
	}
}
func mulvec(a, x, y []float64, m, n int) {
	i, j := 0, 0
	for i = 0; i < m; i++ {
		y[i] = 0
		for j = 0; j < n; j++ {
			y[i] += x[j] * a[i*n+j]
		}

	}
}

func transpose(a, at []float64, m, n int) {
	i, j := 0, 0
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			at[j*m+i] = a[i*n+j]
		}
	}
}

/* A <- A + B */
func accum(a, b []float64, m, n int) {
	i, j := 0, 0
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			a[i*n+j] += b[i*n+j]
		}
	}
}

/* C <- A + B */
func add(a, b, c []float64, n int) {
	j := 0
	for j = 0; j < n; j++ {
		c[j] = a[j] + b[j]
	}
}

/* C <- A - B */
func sub(a, b, c []float64, n int) {
	j := 0
	for j = 0; j < n; j++ {
		c[j] = a[j] - b[j]
	}
}

func negate(a []float64, m, n int) {
	i, j := 0, 0
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			a[i*n+j] = -a[i*n+j]
		}
	}
}

func mat_addeye(a []float64, n int) {
	i := 0
	for i = 0; i < n; i++ {
		a[i*n+i] += 1
	}
}
func cum_addeye(a []float64, n int,f func()float64) {
	i := 0
	for i = 0; i < n; i++ {
		a[i*n+i] += f()
	}
}
func cum_init(a []float64, f func()float64) {
	i := 0
	for i = 0; i < len(a); i++ {
		a[i] += f()
	}
}
/* TinyEKF code ------------------------------------------------------------------- */

type ekf_t struct {
	x []float64 /* state vector */

	P []float64 /* prediction error covariance */
	Q []float64 /* process noise covariance */
	R []float64 /* measurement error covariance */

	G []float64 /* Kalman gain; a.k.a. K */

	F []float64 /* Jacobian of process model */
	H []float64 /* Jacobian of measurement model */

	Ht []float64 /* transpose of measurement Jacobian */
	Ft []float64 /* transpose of process Jacobian */
	Pp []float64 /* P, post-prediction, pre-update */

	fx []float64 /* output of user defined f() state-transition function */
	hx []float64 /* output of user defined h() measurement function */

	/* temporary storage */
	tmp0 []float64
	tmp1 []float64
	tmp2 []float64
	tmp3 []float64
	tmp4 []float64
	tmp5 []float64
}

func ekf_init(ekf *ekf_t, n, m int) {
	ekf.x=_malloc(n,1)
	ekf.P=_malloc(n,n)
	ekf.Q=_malloc(n,n)
	ekf.R=_malloc(m,m)
	ekf.G=_malloc(n,m)
	ekf.F=_malloc(n,n)
	ekf.H=_malloc(m,n)

	ekf.Ht=_malloc(n,m)
	ekf.Ft=_malloc(n,n)
	ekf.Pp=_malloc(n,n)

	ekf.tmp0=_malloc(n,n)
	ekf.tmp1=_malloc(n,m)
	ekf.tmp2=_malloc(m,n)
	ekf.tmp3=_malloc(m,m)
	ekf.tmp4=_malloc(m,m)
	ekf.tmp5=_malloc(m,1)

	ekf.fx=_malloc(n,1)
	ekf.hx=_malloc(m,1)
	/* zero-out matrices */
	zeros(ekf.P, n, n)
	zeros(ekf.Q, n, n)
	zeros(ekf.R, m, m)
	zeros(ekf.G, n, m)
	zeros(ekf.F, n, n)
	zeros(ekf.H, m, n)
}

func ekf_step(ekf *ekf_t, z []float64, n, m int) int {
	/* P_k = F_{k-1} P_{k-1} F^T_{k-1} + Q_{k-1} */
	mulmat(ekf.F, ekf.P, ekf.tmp0, n, n, n)
	transpose(ekf.F, ekf.Ft, n, n)
	mulmat(ekf.tmp0, ekf.Ft, ekf.Pp, n, n, n)
	accum(ekf.Pp, ekf.Q, n, n)

	/* G_k = P_k H^T_k (H_k P_k H^T_k + R)^{-1} */
	transpose(ekf.H, ekf.Ht, m, n)
	mulmat(ekf.Pp, ekf.Ht, ekf.tmp1, n, n, m)
	mulmat(ekf.H, ekf.Pp, ekf.tmp2, m, n, n)
	mulmat(ekf.tmp2, ekf.Ht, ekf.tmp3, m, n, m)
	accum(ekf.tmp3, ekf.R, m, m)
	if cholsl(ekf.tmp3, ekf.tmp4, ekf.tmp5, m) != 0 {
		return 1
	}
	mulmat(ekf.tmp1, ekf.tmp4, ekf.G, n, m, m)

	/* \hat{x}_k = \hat{x_k} + G_k(z_k - h(\hat{x}_k)) */
	sub(z, ekf.hx, ekf.tmp5, m)
	mulvec(ekf.G, ekf.tmp5, ekf.tmp2, n, m)
	add(ekf.fx, ekf.tmp2, ekf.x, n)

	/* P_k = (I - G_k H_k) P_k */
	mulmat(ekf.G, ekf.H, ekf.tmp0, n, m, n)
	negate(ekf.tmp0, n, n)
	mat_addeye(ekf.tmp0, n)
	mulmat(ekf.tmp0, ekf.Pp, ekf.P, n, n, n)

	/* success */
	return 0
}

//形参mu为正态分布均值，sigma为正态分布方差，x为所取的自变量

func Normal( mu, sigma, x float64)float64 {
	result:=0.0 //结果变量
	result=math.Exp((-1)*(x-mu)*(x-mu)/(2*sigma*sigma))/(math.Sqrt(2*math.Pi)*sigma)
	return result
}
