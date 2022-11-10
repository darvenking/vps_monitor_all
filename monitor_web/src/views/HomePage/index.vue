<template>
  <div class="home-page">
    <div class="table-title">
      <div class="table-title-left">
        服务器商家：
        <el-select v-model="id" placeholder="请选择" @change="changeSell" :clearable="true">
          <el-option
            v-for="item in sellers"
            :key="item.ID"
            :label="item.SellerName"
            :value="item.ID">
          </el-option>
        </el-select>

        状态：
        <el-select v-model="status" placeholder="请选择" @change="changeSell" :clearable="true">
          <el-option
            v-for="item in statusFilter"
            :key="item.value"
            :label="item.name"
            :value="item.value">
          </el-option>
        </el-select>
        <el-button type="primary" @click="changeSell" :loading="loading" style="margin-left: 10px">{{
            btnName
          }}
        </el-button>
      </div>
      <div class="table-title-right">
        <el-button type="primary" @click="()=>{dialogFormVisible = true}" style="margin-right: 10px">提交网址
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
          <el-table-column prop="Name" label="名字" width="180"></el-table-column>
          <el-table-column prop="URL" label="链接" width="450"></el-table-column>
          <el-table-column prop="Price" label="价格" width="180"></el-table-column>
          <el-table-column prop="Stock" label="是否有货" width="180">
            <template slot-scope="scope">
              <el-tag
                :type="scope.row.Stock ? 'success' : 'danger'"
                disable-transitions>{{ scope.row.Stock ? '有货' : '缺货' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="100">
            <template slot-scope="scope">
              <template v-if="scope.row.Stock">
                <el-button @click="open_page(scope.row.URL)" type="success" size="small">购买</el-button>
              </template>
              <template v-else>
                <el-button type="danger" size="small" disabled>购买</el-button>
              </template>
            </template>
          </el-table-column>
        </el-table>
      </template>
    </el-skeleton>

    <el-dialog title="提交商品" :visible.sync="dialogFormVisible" :close-on-click-modal="false">
      <el-form :model="form" :rules="rules" ref="ruleForm" label-width="120px">
        <el-form-item label="商家名称" prop="name">
          <el-select v-model="form.name" placeholder="请选商家名称">
            <el-option
              v-for="item in sellers"
              :key="item.ID"
              :label="item.SellerName"
              :value="item.SellerName">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="购买链接" prop="url">
          <el-input v-model="form.url" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="商品名" prop="productName">
          <el-input v-model="form.productName" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="价格" prop="price">
          <el-input v-model="form.price" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="()=>{dialogFormVisible = false}">取 消</el-button>
        <el-button type="primary" @click="submit" :loading="subLoading">{{ subBtnName }}</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import { GetPlistApi, GetSellerApi, SubmitApi } from '@/apis/data.api';
import useCounterStore from '@/store/modules/app';
import { useRoute } from '@/router';
import { open_page } from '@/utils/util';

export default {
  data() {
    return {
      loading: true,
      btnName: '加载中',
      subLoading: false,
      subBtnName: '确 定',
      testNum: 0,
      tableData: [],
      store: useCounterStore(),
      useRoute,
      open_page,
      sellers: [],
      statusFilter: [
        {
          name: '缺货',
          value: 1,
        },
        {
          name: '有货',
          value: 2,
        },
      ],
      id: undefined,
      status: undefined,
      dialogFormVisible: false,
      form: {
        name: '',
        url: '',
        price: '',
        productName: '',
      },
      rules: {
        name: [
          { required: true, message: '请选择商家', trigger: 'blur' },
        ],
        url: [
          { required: true, message: '请填写购买地址', trigger: 'blur' },
        ],
        price: [
          { required: true, message: '请填写购买价格', trigger: 'blur' },
        ],
        productName: [
          { required: true, message: '请填写商品名称', trigger: 'blur' },
        ],
      },
    };
  },
  mounted() {
    this.getSellers();
    this.getPlistApi();
  },
  computed: {},
  methods: {
    async getSellers() {
      let res = await GetSellerApi();
      this.sellers = res.data;
    },
    async getPlistApi() {
      this.loading = true;
      this.tableData = [];
      let res = await GetPlistApi({ id: this.id, stock: this.status });
      this.tableData = res.data;
      this.btnName = '刷新';
      this.loading = false;
    },

    changeSell() {
      this.getPlistApi();
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
        let res = await SubmitApi(this.form);
        this.form = {
          name: '',
          url: '',
          price: '',
          productName: '',
        };
        this.subBtnName = '确 定';
        this.dialogFormVisible = false;
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
