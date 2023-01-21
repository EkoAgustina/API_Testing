const request = require('supertest');
const expect = require('chai').expect;
const env = require('../../env.json');
const userData = require('../../test-data/user/testdata_user.json');
var supertest = require('supertest');
var agent     = supertest.agent(env.baseUrl);

const header = {
    'Accept': 'application/json',
    'Content-Type': 'application/json'
}

describe('USER', () =>{

    it('Successfully created user', () => {
        return new Promise((resolve,reject) => {
            agent
            .post('/register')
            .set(header)
            .send(userData)
            .expect(function (res) {
                if (res.statusCode == 200){
                    console.log('Create user, responseCode:'+res.body.responseCode)
                }
                else{
                    reject(new Error('Create user response code not expected, responseCode:'+res.body.responseCode))
                }

            })
            .end((err,ress)=>{
                if(err){
                    reject(new Error('There an error: '+res.error))
                }
                else{
                    resolve(ress)
                }
            })
        })
    })

    it('Login successfully', () => {
        return new Promise((resolve,reject) => {
            const reqBody = {
                'username': userData.username,
                'password': '123'
            }
            agent
            .post('/login')
            .set(header)
            .send(reqBody)
            .expect(function (res) {
                if (res.statusCode == 200){
                    console.log('Login, responseCode:'+res.body.responseCode)
                }
                else{
                    reject(new Error('Login response Code not expected, responseCode:'+res.body.responseCode))
                }

            })
            .end((err,ress)=>{
                if(err){
                    reject(new Error('There an error: '+res.error))
                }
                else{
                    resolve(ress)
                }
            })
        })
    })
})

module.exports = {
    agent
}