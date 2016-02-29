// API Server utilities
API = {
   _uri: '/api/json',
};

API.get = function(path, data) {
   var form = new FormData();
   _.each(data,function(val, key){
      form.append(key, val)
   })
   fetch(_uri + path, {
      credentials: 'include',
      method: 'post',
      body: form
   })
}
