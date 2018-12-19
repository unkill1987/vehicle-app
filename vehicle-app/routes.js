//SPDX-License-Identifier: Apache-2.0
 
var vehicle = require('./controller.js');

module.exports = function(app){ 

  app.post('/history/:id', function(req, res){
    vehicle.carHistory(req, res);
  });
  app.post('/all', function(req, res){
    vehicle.Allcars(req, res);
  });
  app.post('/init_car/:record_car', function(req, res){
    vehicle.Newcar(req, res);
  });
  app.post('/change_car/:record_change', function(req, res){
    vehicle.Change(req, res);
  });
  app.post('/repair_car/:record_repair', function(req, res){
    vehicle.Repair(req, res);
  });
  app.post('/accident_car/:record_accident', function(req, res){
    vehicle.Accident(req, res);
  });
  
}
