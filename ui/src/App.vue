<template>
    <div style="width: 85%;margin: auto">
    <el-table :data="users" >
    <el-table-column prop="id" label="ID" width="60" />
    <el-table-column prop="name" label="Name" width="200" />
    <el-table-column prop="email" label="Email" width="200" />
    <el-table-column prop="phone" label="Phone" width="180" />
    <el-table-column align="right">
      <template #header>
        <el-input v-debounce:300ms="getUsers" v-model="search" size="small" placeholder="Type to search" />
      </template>
      <template #default="scope">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)"
          >Edit</el-button
        >
        <el-popconfirm
          confirm-button-text="Yes"
          cancel-button-text="No, Thanks"
          :icon="InfoFilled"
          icon-color="#626AEF"
          title="Are you sure to delete this?"
          @confirm="confirmEvent"
          @cancel="cancelEvent"
        >
          <template #reference>
            <el-button
                size="small"
                type="danger"
                @click="handleDelete(scope.$index, scope.row)"
                >Delete</el-button>
          </template>
        </el-popconfirm>
        
      </template>
    </el-table-column>
  </el-table>
  <el-button size="small" @click="handleCreate()"
          >add New user</el-button
        >
  <el-pagination background layout="prev, pager, next" 
    :total="pagination.total"  
    @prev-click="prevPage" 
    @next-click="nextPage" 
    @current-change="handleCurrentChange"
    :current-page="pagination.page"
    :page-size="pagination.limit"  
  />
  
  <!-- user modal -->
  <el-dialog v-model="dialogFormVisible" title="Edit user info">
    <el-form :model="selected_user">
      <el-form-item label="Name" :label-width="formLabelWidth">
        <el-input v-model="selected_user.name" autocomplete="off" />
      </el-form-item>
      <el-form-item label="Email" :label-width="formLabelWidth">
        <el-input v-model="selected_user.email" autocomplete="off" />
      </el-form-item>
      <el-form-item label="Phone" :label-width="formLabelWidth">
        <el-input v-model="selected_user.phone" autocomplete="off" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogFormVisible = false">Cancel</el-button>
        <el-button type="primary" @click="confirmEvent"
          >Confirm</el-button
        >
      </span>
    </template>
  </el-dialog>

  </div>


</template>

<script>

import { ElNotification } from 'element-plus'
export default {
  name: 'App',
  components: {
    
  },
  data() {
    return {
      search : '',
      users : [],
      selected_user : {},
      dialogFormVisible : false,
      action : '',
      actions:{
        'create' : this.createUser,
        'update' : this.updateUser,
        'delete' : this.deleteUser
      },
      pagination : {
        page : 1,
        limit : 12,
        total : 0
      },
    }
  },
  created() {
    this.getUsers();
  },
  methods: {
    getUsers() {
      fetch('http://localhost:8086/api/v1/users/?search='+this.search+'&page='+this.pagination.page+'&limit='+this.pagination.limit)
        .then(response => response.json())
        .then(json => {
          this.users = json.data;
          this.pagination = json.pagination;
        })
    },
    createUser() {
      fetch('http://localhost:8086/api/v1/users/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(this.selected_user)
      })
        
        .then(res => {
          console.log(res.json);
          // if status code is 200
          
          if (res.ok) {
            this.dialogFormVisible = false;
            this.getUsers();
            ElNotification.success({
              title: 'Success',
              message: 'User created successfully',
              duration: 2000
            })
          } else {
            ElNotification.error({
              title: 'Error',
              message: res.json.message,
              duration: 2000
            })
          }
        })
    },
    deleteUser(id) {
      fetch('http://localhost:8086/api/v1/users/'+id, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then(response => response.json())
        .then(json => {
          console.log(json);
          this.getUsers();
          ElNotification({
            title: 'Success',
            message: 'User deleted successfully',
            type: 'success',
          })
        })
    },
    updateUser(id) {
      fetch('http://localhost:8086/api/v1/users/'+id, {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(this.selected_user)
      })
        .then(response => response.json())
        .then(json => {
          console.log(json);
          this.getUsers();
          ElNotification({
            title: 'Success',
            message: 'User updated successfully',
            type: 'success',
          })
          this.dialogFormVisible = false
        })
    },
    prevPage(val) {
      this.pagination.page = val;
    },
    nextPage(val) {
      this.pagination.page = val;
    },
    handleCurrentChange(val) {
      this.pagination.page = val;
      this.getUsers();
    },
    handleCreate(){
      this.action = 'create'
      this.selected_user = {
        name : null,
        email : null,
        phone: null
      }
      this.dialogFormVisible = true;
    },
    handleEdit(_index, row) {
      this.selected_user = row;
      this.dialogFormVisible = true;
      this.action= 'update'
    },
    handleDelete(_index, row){
      this.selected_user = row;
      this.action = 'delete'
    },
  
    cancelEvent() {
      console.log('cancel');
      this.selected_user = {};
      this.action = '';
    },
    confirmEvent() {
      this.actions[this.action](this.selected_user.id);
    }

  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
