<template>
  <div class="home-page">
    <template v-if="!loginStatus">
      <el-input v-model="token" placeholder="请输入token"
                style="width: 200px;display: inline-block;margin-right: 10px"></el-input>
      <el-button type="primary" @click="loginBtn">登录</el-button>
    </template>
    <template v-else>
      <div class="table-title">
        <div class="table-title-left">
          状态：
          <el-select v-model="page.status" placeholder="请选择" @change="changeStatus" :clearable="true">
            <el-option
              v-for="item in statusFilter"
              :key="item.value"
              :label="item.name"
              :value="item.value">
            </el-option>
          </el-select>
          <el-button type="primary" @click="changeStatus" :loading="loading" style="margin-left: 10px">{{
              btnName
            }}
          </el-button>
        </div>
        <div class="table-title-right">
          <el-button type="primary" @click="logout" style="margin-right: 10px">退出登录
          </el-button>
        </div>
      </div>

      <el-skeleton
        :loading="loading"
        animated
        :throttle="500"
      >
        <template slot="template">
          <div style="padding: 14px;">
            <el-skeleton-item variant="h1"/>
            <el-skeleton-item variant="h1"/>
          </div>
        </template>
        <template>
          <el-table :data="tableData" style="width: 100%">
            <el-table-column prop="ID" label="id" width="100"></el-table-column>
            <el-table-column prop="URL" label="链接"></el-table-column>
            <el-table-column prop="Status" label="状态">
              <template slot-scope="scope">
                <el-tag
                  :type="scope.row.Status === 2 ? 'success' : 'danger'"
                  disable-transitions>{{ scope.row.Status === 2 ? '已处理' : '未处理' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="CreatedAt" label="提交日期" width="150">
              <template slot-scope="scope">
                <span>
                  {{moment(scope.row.CreatedAt).format("YYYY-MM-DD HH:mm")}}
                </span>
              </template>
            </el-table-column>
            <el-table-column
              label="操作"
              width="200">
              <template slot-scope="scope">
                <template v-if="scope.row.Status === 1">
                  <el-button @click="open_page(scope.row.URL)" type="warning" size="small">浏览</el-button>
                  <el-button @click="auditBtnFunc(scope.row)" type="success" size="small">处理</el-button>
                </template>
                <template v-else>
                  <el-button type="danger" size="small" disabled>已处理</el-button>
                </template>
              </template>
            </el-table-column>
          </el-table>
          <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :page-sizes="[20, 50, 100]"
            :current-page.sync="page.page"
            :page-size="page.size"
            :total="total"
            layout="sizes, prev, pager, next"
            style="margin-top: 10px"
          >
          </el-pagination>
        </template>
      </el-skeleton>

      <el-dialog title="审核" :visible.sync="dialogFormVisible" :close-on-click-modal="false">
        <el-form :model="form" :rules="rules" ref="ruleForm" label-width="120px">
          <el-form-item label="名字选择器" prop="nameFlag">
            <el-input v-model="form.nameFlag" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="价格选择器" prop="priceFlag">
            <el-input v-model="form.priceFlag" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="无库存关键字">
            <el-input v-model="form.noStockFlag" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item label="Cookies">
            <el-input v-model="form.cookies" autocomplete="off"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="()=>{dialogFormVisible = false}">取 消</el-button>
          <el-button type="primary" @click="submit" :loading="subLoading">{{ subBtnName }}</el-button>
        </div>
      </el-dialog>
    </template>
  </div>
</template>
<script>
import { Audit, AuditList } from '@/apis/data.api';
import userStore from '@/store/modules/app';
import { useRoute } from '@/router';
import { open_page } from '@/utils/util';
import moment from "moment"

export default {
  data() {
    return {
      loading: true,
      btnName: '加载中',
      subLoading: false,
      subBtnName: '确 定',
      store: userStore(),
      moment,
      useRoute,
      open_page,
      token: '',
      loginStatus: false,
      total: 0,
      tableData: [],
      page: {
        page: 1,
        size: 10,
        status: undefined,
      },
      statusFilter: [
        {
          name: '未处理',
          value: 1,
        },
        {
          name: '已处理',
          value: 2,
        },
      ],
      dialogFormVisible: false,
      form: {
        id: 0,
        nameFlag: '',
        priceFlag: '',
        cookies: '',
        noStockFlag: '',
      },
      rules: {
        nameFlag: [
          { required: true, message: '请填写名字选择器', trigger: 'blur' },
        ],
        priceFlag: [
          { required: true, message: '请填写价格选择器', trigger: 'blur' },
        ],
      },
    };
  },
  mounted() {
    this.token = localStorage.getItem('secret');
    if (!this.token) {
      this.loginStatus = false;
    }
    this.loginStatus = true;
    this.getAuditList();
  },
  computed: {},
  methods: {
    handleSizeChange(v) {
      this.page.size = v;
    },
    handleCurrentChange(v) {
      this.page.page = v;
    },
    loginBtn() {
      localStorage.setItem('secret', this.token.toString());
      this.loginStatus = true;
      this.getAuditList();
    },
    logout() {
      this.loginStatus = false;
      this.token = '';
      localStorage.setItem('secret', this.token);
    },
    async getAuditList() {
      this.loading = true;
      this.btnName= '加载中';
      let res = await AuditList(this.page);
      if (res.code === 403) {
        this.$notify({
          type:"error",
          title: "错误",
          message: "未登录或授权码错误",
        })
        this.logout();
        return;
      }
      this.tableData = res.data.data;
      this.total = res.data.total;
      this.loading = false;
      this.btnName= '确 定';
    },
    auditBtnFunc(row){
      this.form.id = row.ID;
      console.log("row==>",row);
      console.log("form==>",this.form);
      this.dialogFormVisible = true;
    },
    changeStatus() {
      this.getAuditList(this.page);
    },
    submit() {
      this.subLoading = true;
      this.subBtnName = '提交中';
      this.$refs.ruleForm.validate(async (valid) => {
        if (!valid) {
          console.log('error submit!!');
          this.subLoading = false;
          this.subBtnName = '确 定';
          return false;
        }
        let res = await Audit(this.form);
        this.form = {
          id: undefined,
          nameFlag: '',
          priceFlag: '',
          cookies: '',
          noStockFlag: '',
        };
        this.subLoading = false;
        this.subBtnName = '确 定';
        this.dialogFormVisible = false;
        this.$notify({
          type:"success",
          title: "成功",
          message: "审核成功",
        })
        this.getAuditList();
      });
    },
  },
  components: {},
};
</script>
<style lang="scss" scoped>
.home-page {
  .test-text {
    color: v-bind(textColor);
  }

  .table-title {
    display: flex;
    margin: 20px 0;

    .table-title-right {
      margin-left: auto;
    }
  }
}
</style>
