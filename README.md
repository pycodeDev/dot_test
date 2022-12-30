# dot_test
 Dot Test

Test DoT

Run Project <br/>
<code>go run dot.go</code> <br/>

Credential. <br/>
url = <code>http://127.0.0.1:9003/v1</code><br/>
Api Key = <code>6c1118aa11d0d1944adf85aa86602b5ee47da646</code><br/>
Email = anang@gmail.com<br/>
Pass = 12345<br/>
Nama Database = db_dot_test<br/>

After Running, There is 9 Routes :<br/>
1. <code>/user/login</code>
2. <code>/user/update</code>
3. <code>/buku/insert</code>
4. <code>/buku/update</code>
5. <code>/buku/delete</code>
6. <code>/buku/list</code>
7. <code>/buku/:id</code>
7. <code>/peminjaman/create</code>
7. <code>/peminjaman/list</code>
<br/>

untuk config mysql ada di xconfig, jika tidak menggunakan password, silahkan hilangkan di bagian <code>MYSQL_PASS_WRITE</code>.

untuk route <b>2</b> hingga <b>9</b>
Menggunakan Middleware di <code>Authorization</code> <br/>
Ambil token <b>JWT</b> setelah Login.
<br/>

Link Dokumentasi = <a href="https://documenter.getpostman.com/view/21493225/2s8Z6yXYTi" target="_blank">Dokumentasi</a><br/>
Link Screen Recorder Doc APi = <a href="https://www.loom.com/share/98b8a819ad35467ab9bca867049a9723" target="_blank">Cek</a><br/>
Link Screen Recorder Penjelasan Template Base Project = <a href="https://www.loom.com/share/e359357b03c641d3a9bd938e2f17d9a2" target="_blank">Cek</a><br/>
Link Database = <a href="https://ideone.com/BHs90C" target ="_blank">Get</a><br/>

😁 Thanks.<br/>
I Use <a href="https://github.com/gofiber/fiber">⚡️Fiber</a> Framework.
