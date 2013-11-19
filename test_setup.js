var nano = require('nano')('http://garren:password@localhost:5984');

var docs = [
  {
  "_id": "16b57b197b487766dc2973e9806bd0c7",
  "user_id": "2",
  "user_name": "Bill Pearl",
  "organisation_id": "1",
  "category": "assessment_completed",
  "info": {
    "name": "Third assessment"
  },
  "link": "/assessments/99/show_answers?learner_id=2",
  "updated_at": "2013-10-22T15:08:51Z",
  "created_at": "2013-10-22T09:04:37Z",
  "type": "ActivityItem"
},

{
  "_id": "16b57b197b487766dc2973e9806be3ef",
  "user_id": "2",
  "user_name": "Bill Pearl",
  "organisation_id": "1",
  "category": "assessment_completed",
  "info": {
    "name": "Second Assessment"
  },
  "link": "/assessments/98/show_answers?learner_id=2",
  "updated_at": "2013-10-28T15:08:51Z",
  "created_at": "2013-10-28T09:04:06Z",
  "type": "ActivityItem"
},

{
  "_id": "16b57b197b487766dc2973e9806bf62a",
  "_rev": "2-5d92e08d4de902392ede93c6b2d2f9b3",
  "user_id": "2",
  "user_name": "Bill Pearl",
  "organisation_id": "1",
  "category": "assessment_reopened",
  "info": {
    "name": "Second Assessment"
  },
  "link": "/assessments/98/show_answers?learner_id=2",
  "updated_at": "2013-10-20T15:08:52Z",
  "created_at": "2013-10-20T09:03:28Z",
  "type": "ActivityItem"
}
];


// kids don't design views like this at home!
var ddoc = {
  "_id": "_design/ActivityItem",
  "language": "javascript",
  "views": {
    "testView": {
      "map": "function(doc) {\n                  if ((doc['type'] == 'ActivityItem') && (doc['organisation_id'] !== null) && (doc['created_at'] !== null)) {\n                    emit([doc['organisation_id'], doc['updated_at']], 1);\n                  }\n                }\n",
      "reduce": "_sum"
    }
  }
};

nano.db.destroy('nano-go-tests', function() {

  nano.db.create('nano-go-tests', function () {
    var db = nano.use('nano-go-tests');
    
    docs.forEach(function(doc) {
      db.insert(doc, doc._id);
    });

    db.insert(ddoc, ddoc._id);
  });
});
