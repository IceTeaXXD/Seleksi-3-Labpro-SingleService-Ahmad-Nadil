import { PrismaClient } from '@prisma/client';
import express, { Request, Response } from 'express';

const prisma = new PrismaClient();
const app = express();

// Endpoint CRUD Barang
app.get('/products', async (req: Request, res: Response) => {
  const products = await prisma.product.findMany();
  res.json(products);
});

app.post('/products', async (req: Request, res: Response) => {
  // Validasi admin token
  // Ambil data barang dari body request
  // Lakukan validasi data barang
  // Simpan data barang ke basis data menggunakan Prisma Client
  // Kirim response berhasil
});

// Implementasikan endpoint CRUD Perusahaan

// Implementasikan endpoint Login Admin

app.listen(3000, () => {
  console.log('Server is running on port 3000');
});
