# งานพัฒนาการจำลองการจัดการโปรเซส 2
### จงพัฒนาโปรแกรมการจำลองโปรเซส โดยให้มีคุณสมบัติดังนี้
- มี CPU จำนวน 2 ชุด
- มี ready queue แบบ Multilevel 3 ชุด โดยให้มีการจัดลำดับ priority สำคัญกว่าทำงาน 3 โปรเซส แล้วจึงให้ priority สำคัญน้อยกว่าทำงาน 1 โปรเซส
- มี I/O queue  4 ชุด

### และมีชุดคำสั่งดังนี้
- สร้างโปรเซส พร้อมกำหนด priority
- จบโปรเซส
- โปรเซสหมดเวลาใช้หน่วยประมวลผลกลาง
- โปรเซสเข้า I/O คิว
- โปรเซสออกจาก I/O คิว