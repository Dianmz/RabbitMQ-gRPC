const express=require('express');
const app=new express();

app.use(express.json({limit:'5mb', extended:true}));

app.post('/', async (req,res)=>{
    try{
        //res.json(result.ops[0]);
        res.json(req.body);
    }catch(err){
        console.log(err);
        res.status(500).json({'message':'failed'});
    }
});

app.get('/', (req, res) => {
    res.json({message: 'OK'})
});

app.listen(5000);