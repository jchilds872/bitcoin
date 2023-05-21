d={}

qq,tt,ii,oo,bb=0,0,0,0,0

for l in open('data_test.txt','r'):
   v=l.split(" ")
   b,t,i,o=int(v[0]),int(v[1]),int(v[2]),int(v[3])
   tt+=t
   ii+=i
   oo+=o
   for x in range(4+i,4+i+o):
    vv=float(v[x])
    if vv>0: d[v[x+o][-16:]]=(b,vv); bb+=vv
   for x in range(4,4+i): d.pop(v[x][-16:], None)
   
print(len(d))
